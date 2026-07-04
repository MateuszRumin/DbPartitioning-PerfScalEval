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

func main() {
	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()
	e := 0

	start := time.Now()

	for _, query := range PostInsertFull {

		rows, err := db.Query(query)
		if err != nil {
			e++
		}
		defer rows.Close()

	}

	stop := time.Now()

	db2, err := slc()
	if err != nil {
		return
	}
	defer db2.Close()

	db2.Exec(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('%s','%s','%s')",
		"Insertds rekord Post", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

}

func slc() (*sql.DB, error) {

	user := "root"
	password := ""
	host := "192.168.50.3"
	port := "3306"
	database := "logdb"
	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	// Połączenie z bazą danych
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
