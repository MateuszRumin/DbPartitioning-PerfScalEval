package main

import (
	crand "crypto/rand"
	"database/sql"
	"encoding/binary"
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

func newWorkerRand() *rand.Rand {
	var b [16]byte
	_, _ = crand.Read(b[:])

	s1 := binary.LittleEndian.Uint64(b[:8])
	s2 := binary.LittleEndian.Uint64(b[8:])

	return rand.New(rand.NewPCG(s1, s2))
}

func wantConnection(id int, r *rand.Rand, wg *sqlgen.WorkerGenerator, idp int, idu int) {
	fmt.Println("Worker:", id)

	db, err := setConnection()
	if err != nil {
		log.Printf("[worker %d] DB error: %v", id, err)
		return
	}
	defer db.Close()

	deadline := time.Now().Add(10 * time.Minute)

	for time.Now().Before(deadline) {

		query := wg.GenerateRandomQuery(r, idp, idu)
		if query == "" {
			continue
		}

		err := executeQuery(db, query)
		if err != nil {
			log.Printf("[worker %d] query error: %v", id, err)
			continue
		}
	}
}

func multiThreadConnection() {
	idp, idu, err := sqlgen.GenerateInitValues()
	if err != nil {
		log.Println("Błąd generacji zapytania")
		return
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			r := newWorkerRand()
			wg := sqlgen.NewWorkerGenerator(r)
			wantConnection(id, r, wg, idp, idu)
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

	count := 0
	for rows.Next() {
		count++
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows iteration error: %w", err)
	}

	//fmt.Println("i")

	return nil
}
