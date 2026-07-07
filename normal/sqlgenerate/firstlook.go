package sqlgenerate

import (
	"database/sql"
	"fmt"

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

func GenerateInitValues() (int, int, error) {
	db, err := setConnection()
	if err != nil {
		return 0, 0, err
	}
	defer db.Close()

	var postsCount int
	err = db.QueryRow("SELECT COUNT(id) FROM posts").Scan(&postsCount)
	if err != nil {
		return 0, 0, err
	}

	var usersCount int
	err = db.QueryRow("SELECT COUNT(id) FROM users").Scan(&usersCount)
	if err != nil {
		return 0, 0, err
	}

	return postsCount, usersCount, nil
}
