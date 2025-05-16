package dbconects

import (
	"database/sql"
	"fmt"
	"perfscaleval/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func SetConnections() {
	var wg sync.WaitGroup
	config.Connections = make([]*sql.DB, config.NumConnections)
	var mu sync.Mutex

	errors := make([]error, config.NumConnections)
	for i := 0; i < config.NumConnections; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			conn, err := setConnection()
			if err != nil {
				errors[i] = err
				fmt.Printf("Błąd podczas łączenia %d: %v\n", i, err)
				return
			}

			if err := conn.Ping(); err != nil {
				errors[i] = err
				fmt.Printf("Błąd podczas pingowania %d: %v\n", i, err)
				conn.Close()
				return
			}

			_, execErr := conn.Exec("SHOW PROCESSLIST;")
			if execErr != nil {
				errors[i] = execErr
				fmt.Printf("Błąd podczas wykonywania zapytania %d: %v\n", i, execErr)
				conn.Close()
				return
			}

			mu.Lock()
			config.Connections[i] = conn
			mu.Unlock()

			fmt.Printf("Połączono %d\n", i+1)
		}(i)
	}
	wg.Wait()

	// Sprawdzenie, czy wszystkie połączenia zostały poprawnie utworzone
	successCount := 0
	for _, err := range errors {
		if err == nil {
			successCount++
		}
	}

	if successCount == config.NumConnections {
		config.ConectionWorking = true
	}

	fmt.Printf("Utworzono poprawnie %d/%d połączeń.\n", successCount, config.NumConnections)
}
