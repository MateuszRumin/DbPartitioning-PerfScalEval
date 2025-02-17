package dbconects

import (
	"database/sql"
	"fmt"
	"sync"

	"perfscaleval/config"
)

// SetConnection jest zaimplementowane w innym pliku i zwraca *sql.DB, error

func SetConnections() {
	var wg sync.WaitGroup
	wg.Add(config.NumConnections)

	config.Connections = make([]*sql.DB, config.NumConnections)
	var mu sync.Mutex

	for i := 0; i < config.NumConnections; i++ {
		i := i
		go func() {
			defer wg.Done()

			conn, err := SetConnection()
			if err != nil {
				fmt.Printf("Błąd podczas łączenia %d: %v\n", i, err)
				return
			}
			conn.Exec("Show processlist;")
			fmt.Printf("połączono %d: \n", i+1)

			mu.Lock()
			config.Connections[i] = conn
			mu.Unlock()
		}()
	}

	wg.Wait()

}
