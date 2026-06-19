package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func setConnection() (*sql.DB, error) {

	user := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "testdb"
	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	// Połączenie z bazą danych
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func wantConnection(id int, db *sql.DB, simpleSelect []string) {
	fmt.Println("Worker: ", id)
	db, err := setConnection()

	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	} else {
		for i := 0; i < 30; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			randomIndex := r.Intn(len(simpleSelect))
			fmt.Println("Worker: ", id, " wykonuje zapytanie: ", simpleSelect[randomIndex])
			err := executeQuery(db, simpleSelect[randomIndex])
			if err != nil {
				log.Printf("[worker %d] błąd: %v", id, err)
				continue
			}

		}

	}
	defer db.Close()
	fmt.Printf("Worker: %d executed Task \n", id)
}

func multiThreadConnection(db *sql.DB) {
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			wantConnection(id, db, simpleSelect)
		}(i)
	}

	wg.Wait()

}

func executeQuery(db *sql.DB, query string) error {
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows iteration error: %w", err)
	}

	return nil
}
