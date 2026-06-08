package main

import (
	"database/sql"
	"fmt"
	"log"

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

func wantConnection(db *sql.DB) {
	db, err := setConnection()
	//err = db.Ping()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	} else {
		fmt.Println("Połączenie z bazą danych działa poprawnie.")

		for _, query := range joinAnalize {
			executeQuery(db, query)
		}

	}
	defer db.Close()

}

func executeQuery(db *sql.DB, query string) {
	fmt.Printf("Wykonuję zapytanie: %s\n", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Nie udało się wykonać zapytania: %v", err)
	}
	defer rows.Close()

	//for rows.Next() {
	//var result int
	//if err := rows.Scan(&result); err != nil {
	//	log.Fatalf("Nie udało się zeskanować wyniku: %v", err)
	//}
	//fmt.Printf("Wynik zapytania: %d\n", result)
	//}

	//if err := rows.Err(); err != nil {
	//log.Fatalf("Błąd podczas iteracji po wynikach: %v", err)
	//}
}
