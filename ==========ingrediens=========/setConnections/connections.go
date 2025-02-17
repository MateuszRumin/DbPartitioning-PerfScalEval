package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

// SetConnection jest zaimplementowane w innym pliku i zwraca *sql.DB, error

func runConnections(numConnections int) []*sql.DB {
	var wg sync.WaitGroup
	wg.Add(numConnections)

	connections := make([]*sql.DB, numConnections)
	var mu sync.Mutex

	for i := 0; i < numConnections; i++ {
		i := i
		go func() {
			defer wg.Done()

			conn, err := SetConnection()
			if err != nil {
				fmt.Printf("Błąd podczas łączenia %d: %v\n", i, err)
				return
			}

			mu.Lock()
			connections[i] = conn
			mu.Unlock()
		}()
	}

	wg.Wait()
	return connections
}

func RunTests(connections []*sql.DB) {
	var wg sync.WaitGroup
	wg.Add(len(connections))

	for i, conn := range connections {
		if conn != nil {
			i := i
			go func() {
				defer wg.Done()
				fmt.Printf("Testowanie połączenia %d...\n", i)
				// Tutaj można wykonać równoczesne zapytania testowe
				query := "SELECT * FROM posts WHERE creation_date BETWEEN '2023-01-01' AND '2023-12-31';"
				_, err := conn.Exec(query)
				if err != nil {
					log.Fatalf("Błąd podczas wykonywania zapytania: %v", err)
				} else {
					log.Println("Krok zaliczony")
				}
			}()
		} else {
			fmt.Printf("Połączenie %d jest nil – nie udało się nawiązać.\n", i)
			wg.Done() // Zmniejszamy licznik, bo nil-owe połączenie nie startuje testu
		}
	}

	wg.Wait()
}

func main() {
	numConnections := 100
	dbConnections := runConnections(numConnections)

	// Testy uruchamiane równolegle
	RunTests(dbConnections)
}
