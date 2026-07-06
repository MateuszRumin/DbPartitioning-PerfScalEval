package main

import (
	"database/sql"
	"fmt"
	"log"
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
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func testEntry() {
	db, err := setConnection()
	//err = db.Ping()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	} else {
		fmt.Println("Połączenie z bazą danych działa poprawnie.")
		exp := Experiment{
			Name:    "NiePartycjonowana30gbTest1Uruchomienie1",
			Queries: simpleSelect,
			Runs:    10,
		}
		results := RunExperiment(db, exp)
		if err := exportCSV("results.csv", results); err != nil {
			log.Fatal(err)
		}
	}
	defer db.Close()

}

func RunExperiment(db *sql.DB, exp Experiment) []QueryResult {
	// WARM-UP
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for _, q := range exp.Queries {
			fmt.Println("i")
			rows, err := db.Query(q)
			if err != nil {
				return nil
			}
			rows.Close()

		}
	}

	var results []QueryResult

	for i := 0; i < exp.Runs; i++ {
		for _, q := range exp.Queries {
			fmt.Println("o")
			start := time.Now()
			rows, err := db.Query(q)
			if err != nil {
				results = append(results, QueryResult{Query: q, Err: err})
				continue
			}

			rowCount := 0

			for rows.Next() {
				rowCount++
				var tmp interface{}
				_ = rows.Scan(&tmp)
			}

			total := time.Since(start)

			rows.Close()

			exRows, parts, _ := Explain(db, q)

			results = append(results, QueryResult{
				Query:       q,
				Duration:    total,
				Rows:        rowCount,
				ExplainRows: exRows,
				Partitions:  parts,
				Err:         nil,
			})
		}
	}

	return results
}
