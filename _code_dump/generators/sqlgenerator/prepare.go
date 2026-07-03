package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func loadIDs(db *sql.DB, table string) ([]int, error) {
	rows, err := db.Query(
		fmt.Sprintf("SELECT id FROM %s", table),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var ids []int

	for rows.Next() {
		var id int

		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}
