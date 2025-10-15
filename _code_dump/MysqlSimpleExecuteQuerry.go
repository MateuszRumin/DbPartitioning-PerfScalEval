package main

import (
	"bufio"
	"database/sql"
	"log"
	"log"
	"os"

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

func SetSimpleProfilingMysql(db *sql.DB) {

	setprofileing := "set profiling = 1"
	_,err := ExecuteQuery(db, setprofileing)
	if err != nil {
		log.Fatalf("Błąd podczas wykonywania zapytania: %v", err)
	}else{
		reader := bufio.NewReader(os.Stdin)
			_, _ = reader.ReadString('\n')
		db.Close()
	}
}