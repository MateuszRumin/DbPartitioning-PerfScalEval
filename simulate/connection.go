package main

import (
	"database/sql"
	"fmt"
	"log"
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

func executeQuery(db *sql.DB, query string) (err error) {
	// fmt.Printf("Wykonuję zapytanie: %s\n", query)
	rows, err := db.Query(query)
	if err != nil {
		//log.Printf("Zapytanie, które spowodowało błąd: %s\n", query)
		log.Printf("Nie udało się wykonać zapytania: \n")
		return err

	}

	defer rows.Close()
	return nil

}

func executeSimpleQuery(db *sql.DB, query string) (int, error) {
	var result int

	err := db.QueryRow(query).Scan(&result)
	if err != nil {
		return 0, err
	}

	return result, nil
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

	return db, nil
}

func executeQuerySelect(db *sql.DB, query string) error {
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
