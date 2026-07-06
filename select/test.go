package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func executeQuery(db *sql.DB, query string) (
	duration time.Duration,
	rowCount int,
	err error,
) {
	start := time.Now()

	rows, err := db.Query(query)
	if err != nil {
		return time.Since(start), 0, err
	}

	columns, err := rows.Columns()
	if err != nil {
		rows.Close()
		return time.Since(start), 0, err
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

			firstRowMeasured = true
		}

		if err := rows.Scan(destinations...); err != nil {
			rows.Close()
			return time.Since(start), rowCount, err
		}

		rowCount++
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		return time.Since(start), rowCount, err
	}

	if err := rows.Close(); err != nil {
		return time.Since(start), rowCount, err
	}

	duration = time.Since(start)

	return duration, rowCount, nil
}

func RunExperiment(db *sql.DB, exp Experiment) []QueryResult {
	// EXPLAIN wykonujemy raz dla każdego zapytania.
	explainCache := make(map[string]QueryResult, len(exp.Queries))

	for _, query := range exp.Queries {
		exRows, partitions, err := Explain(db, query)
		if err != nil {
			explainCache[query] = QueryResult{
				Query: query,
				Err:   fmt.Errorf("EXPLAIN: %w", err),
			}
			continue
		}

		explainCache[query] = QueryResult{
			Query:       query,
			ExplainRows: exRows,
			Partitions:  partitions,
		}
	}

	type aggregate struct {
		totalDuration time.Duration
		totalRows     int64
		successCount  int64
		firstErr      error
	}

	aggregates := make(map[string]*aggregate, len(exp.Queries))

	for _, query := range exp.Queries {
		aggregates[query] = &aggregate{}
	}

	for run := 0; run < exp.Runs; run++ {
		for _, query := range exp.Queries {
			duration, rowCount, err := executeQuery(db, query)

			agg := aggregates[query]

			if err != nil {
				if agg.firstErr == nil {
					agg.firstErr = err
				}
				continue
			}

			agg.totalDuration += duration
			agg.totalRows += int64(rowCount)
			agg.successCount++
		}
	}

	results := make([]QueryResult, 0, len(exp.Queries))

	for _, query := range exp.Queries {
		agg := aggregates[query]
		explain := explainCache[query]

		result := QueryResult{
			Query:       query,
			ExplainRows: explain.ExplainRows,
			Partitions:  explain.Partitions,
		}

		if agg.successCount == 0 {
			if agg.firstErr != nil {
				result.Err = agg.firstErr
			} else {
				result.Err = fmt.Errorf("brak poprawnych wykonań zapytania")
			}

			results = append(results, result)
			continue
		}

		result.Duration = agg.totalDuration / time.Duration(agg.successCount)
		result.Rows = int(agg.totalRows / agg.successCount)

		result.Err = agg.firstErr

		results = append(results, result)
	}

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
