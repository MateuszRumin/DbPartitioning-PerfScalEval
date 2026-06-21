package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"time"

	sqlgen "pkSelectTestGen/sqlgenerate"

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

func wantConnection(id int) {
	fmt.Println("Worker:", id)

	db, err := setConnection()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	}
	defer db.Close()

	ids, err := sqlgen.GenerateInitValues()
	if err != nil {
		log.Println("Błąd generacji zapytania")
		return
	}
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d\n",
		len(ids["badges"]),
		len(ids["users"]),
		len(ids["votes"]),
		len(ids["tags"]),
		len(ids["posts"]),
		len(ids["post_history"]),
		len(ids["post_links"]),
		len(ids["comments"]),
	)

	r := rand.New(rand.NewPCG(uint64(id), uint64(time.Now().UnixNano())))

	deadline := time.Now().Add(10 * time.Minute)

	for time.Now().Before(deadline) {

		query := sqlgen.GenerateRandomQuery(r, ids)
		fmt.Printf("%s\n\n", query)

		start := time.Now()

		err := executeQuery(db, query)

		duration := time.Since(start)

		metrics.Add(duration)

		if err != nil {
			log.Printf("[worker %d] błąd: %v", id, err)
			continue
		}
	}

	//fmt.Printf("Worker: %d wykonano\n", id)
}

func multiThreadConnection() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			wantConnection(id)
		}(i)
	}

	wg.Wait()

	PrintMinuteReport()
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
