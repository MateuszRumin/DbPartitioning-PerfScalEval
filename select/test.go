package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func executeQuery(db *sql.DB, query string) (
	duration time.Duration,
	firstRow time.Duration,
	rowCount int,
	err error,
) {
	start := time.Now()

	rows, err := db.Query(query)
	if err != nil {
		return time.Since(start), 0, 0, err
	}

	columns, err := rows.Columns()
	if err != nil {
		rows.Close()
		return time.Since(start), 0, 0, err
	}

	// Przygotowanie odbiorników dla dowolnej liczby kolumn.
	values := make([]any, len(columns))
	destinations := make([]any, len(columns))

	for i := range values {
		destinations[i] = &values[i]
	}

	firstRowMeasured := false

	for rows.Next() {
		if !firstRowMeasured {
			firstRow = time.Since(start)
			firstRowMeasured = true
		}

		if err := rows.Scan(destinations...); err != nil {
			rows.Close()
			return time.Since(start), firstRow, rowCount, err
		}

		rowCount++
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		return time.Since(start), firstRow, rowCount, err
	}

	if err := rows.Close(); err != nil {
		return time.Since(start), firstRow, rowCount, err
	}

	duration = time.Since(start)

	return duration, firstRow, rowCount, nil
}

func RunExperiment(db *sql.DB, exp Experiment) []QueryResult {

	// EXPLAIN wykonujemy raz, poza właściwym pomiarem.
	explainCache := make(map[string]QueryResult)

	for _, query := range exp.Queries {
		exRows, partitions, err := Explain(db, query)
		if err == nil {
			explainCache[query] = QueryResult{
				ExplainRows: exRows,
				Partitions:  partitions,
			}
		}
	}

	results := make(
		[]QueryResult,
		0,
		exp.Runs*len(exp.Queries),
	)
	start := time.Now()

	// Właściwy pomiar.
	for run := 0; run < exp.Runs; run++ {
		for _, query := range exp.Queries {
			duration, firstRow, rowCount, err :=
				executeQuery(db, query)

			explain := explainCache[query]

			results = append(results, QueryResult{
				Query:       query,
				Duration:    duration,
				FirstRow:    firstRow,
				Rows:        rowCount,
				ExplainRows: explain.ExplainRows,
				Partitions:  explain.Partitions,
				Err:         err,
			})
		}
	}

	stop := time.Now()

	db2, err := slc()
	if err != nil {
		return results
	}
	defer db2.Close()

	db2.Exec(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('%s','%s','%s')",
		"Select posth p", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

	return results
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
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
