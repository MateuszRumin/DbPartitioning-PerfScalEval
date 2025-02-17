package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ExecuteQuery(db *sql.DB, query string) int {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Błąd podczas wykonywania zapytania: %v", err)
		return 0
	} else {
		log.Println("Krok zaliczony")
		return 0
	}

}
