package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func SetConnection() (*sql.DB, error) {

	user := ""
	password := ""
	host := ""
	port := ""
	database := ""
	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	// Połączenie z bazą danych
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func WantConnection() {
	db, err := SetConnection()
	//err = db.Ping()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	}else{
		fmt.Println("Połączenie z bazą danych działa poprawnie.")
		query := "SELECT 1"
		ExecuteQuery(db, query)

	}
	defer db.Close()	

} 
