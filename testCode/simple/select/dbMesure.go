package main

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type QueryResult struct {
	Query       string
	Duration    time.Duration
	FirstRow    time.Duration
	Err         error
	Rows        int
	ExplainRows int
	Partitions  string
}

type Experiment struct {
	Name    string
	Queries []string
	Runs    int
}

func Explain(db *sql.DB, query string) (rowsExamined int, partitions string, err error) {
	q := "EXPLAIN " + query

	res, err := db.Query(q)
	if err != nil {
		return 0, "", err
	}
	defer res.Close()

	var (
		id, selectType, table, parts, joinType, possibleKeys, key sql.NullString
		keyLen, ref, rows, filtered, extra                        sql.NullString
	)

	if res.Next() {
		err := res.Scan(
			&id, &selectType, &table, &parts, &joinType,
			&possibleKeys, &key, &keyLen, &ref, &rows,
			&filtered, &extra,
		)
		if err != nil {
			return 0, "", err
		}

		return atoiSafe(rows.String), parts.String, nil
	}

	return 0, "", nil
}

func atoiSafe(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
