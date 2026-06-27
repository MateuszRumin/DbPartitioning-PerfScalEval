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

func executeQuery(db *sql.DB, query string, id int) (err error) {
	// fmt.Printf("Wykonuję zapytanie: %s\n", query)
	rows, err := db.Query(query)
	if err != nil {
		//log.Printf("Zapytanie, które spowodowało błąd: %s\n", query)
		log.Printf("Nie udało się wykonać zapytania:%s \n", err)
		return err

	}
	fmt.Printf("Succes wątek id: %d \n", id)
	defer rows.Close()
	return nil

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

	return db, nil
}
