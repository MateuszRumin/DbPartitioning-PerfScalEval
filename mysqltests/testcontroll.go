package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := SetConnection()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	}
	defer db.Close()
	fmt.Println("Chose test sets")
	err = db.Ping()
	if err != nil {
		fmt.Println("Połączenie z bazą danych nie działa.")

	} else {
		fmt.Println("Połączenie z bazą danych działa poprawnie.")
		idexSerch(db)
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')
		db.Close()
	}

}

func idexSerch(db *sql.DB) {
	setprofileing := "set profiling = 1"
	ExecuteQuery(db, setprofileing)

}
