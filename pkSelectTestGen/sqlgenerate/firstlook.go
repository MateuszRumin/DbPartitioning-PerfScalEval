package sqlgenerate

import (
	"database/sql"
	"fmt"

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

func GenerateInitValues() (map[string][]int, error) {

	db, err := setConnection()
	if err != nil {
		fmt.Println("Błąd połaczenia z bazą danych")
		return nil, err
	}

	tables := []string{"badges", "users", "votes", "tags", "posts", "post_history", "post_links", "comments"}

	ids := make(map[string][]int)

	for _, table := range tables {
		data, err := loadIDs(db, table)
		if err != nil {
			fmt.Println(table)
			return nil, err
		}
		ids[table] = data
	}

	return ids, nil
}
