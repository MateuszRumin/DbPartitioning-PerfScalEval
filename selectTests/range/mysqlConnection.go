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

type QueryResults struct {
	qtype    string
	end      time.Time
	duration time.Duration
}

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

func slc() (*sql.DB, error) {

	user := "root"
	password := ""
	host := "localhost"
	port := "3306"
	database := "logdb"
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
	var qr []QueryResults

	deadline := time.Now().Add(3 * time.Minute)

	for time.Now().Before(deadline) {

		query := wg.GenerateRandomQuery(r, idp, idu)
		if query == "" {
			continue
		}
		start := time.Now()
		err := executeQuery(db, query)
		if err != nil {
			log.Printf("[worker %d] query error: %v", id, err)
			continue
		}
		stop := time.Now()
		duration := time.Since(start)

		qr = append(qr, QueryResults{
			qtype:    "SELECT", // np. SELECT, INSERT, UPDATE
			end:      stop,
			duration: duration,
		})

	}

	db2, err := slc()
	if err != nil {

		return
	}
	defer db.Close()

	for _, d := range qr {

		db2.Query(fmt.Sprintf("INSERT INTO QueryResults (query_type,timeEnded,duration_ms) VALUES ('%s','%s','%d')", d.qtype, d.end.Format("2006-01-02 15:04:05"), d.duration.Milliseconds()))

	}
}

func multiThreadConnection() {
	idp, idu, err := sqlgen.GenerateInitValues()
	if err != nil {
		log.Println("Błąd generacji zapytania")
		return
	}

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			r := newWorkerRand()
			wg := sqlgen.NewWorkerGenerator(r)
			wantConnection(id, r, wg, idp, idu)
		}(i)
	}

	wg.Wait()

	stop := time.Now()

	db, err := slc()
	if err != nil {

		return
	}
	defer db.Close()
	db.Query(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('%s','%s','%s')", "Select Real query", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

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
