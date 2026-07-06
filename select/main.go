package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	testEntry()
}

func testEntry() {
	db, err := setConnection()
	//err = db.Ping()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	} else {
		fmt.Println("Połączenie z bazą danych działa poprawnie.")
		exp := Experiment{
			Name:    "postHistory",
			Queries: Posts,
			Runs:    2,
		}

		results := RunExperiment(db, exp)

		db2, err := slc()
		if err != nil {
			return
		}
		defer db2.Close()

		if err := exportCSV("results.csv", results); err != nil {
			log.Fatal(err)
		}
	}
	defer db.Close()

}

func setConnection() (*sql.DB, error) {

	user := "root"
	password := ""
	host := "192.168.50.3"
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

func exportCSV(path string, results []QueryResult) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{
		"duration_ns",
		"rows",
		"explain_rows",
		"partitions",
	})

	for _, r := range results {
		w.Write([]string{
			strconv.FormatInt(r.Duration.Nanoseconds(), 10),
			strconv.Itoa(r.Rows),
			strconv.Itoa(r.ExplainRows),
			r.Partitions,
		})
	}

	return nil
}
