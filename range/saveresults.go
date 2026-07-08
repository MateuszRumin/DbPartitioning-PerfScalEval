package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func saveQueryResults(
	db *sql.DB,
	resultsByWorker [][]QueryResults,
) error {
	if db == nil {
		return fmt.Errorf("save query results: nil database connection")
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin result transaction: %w", err)
	}

	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	const batchSize = 1000

	batch := make([]QueryResults, 0, batchSize)
	batchNumber := 0

	for _, workerResults := range resultsByWorker {
		for _, result := range workerResults {
			batch = append(batch, result)

			if len(batch) >= batchSize {
				batchNumber++

				if err := insertQueryResultsBatch(tx, batch); err != nil {
					return fmt.Errorf(
						"insert batch %d: %w",
						batchNumber,
						err,
					)
				}

				batch = batch[:0]
			}
		}
	}

	if len(batch) > 0 {
		batchNumber++

		if err := insertQueryResultsBatch(tx, batch); err != nil {
			return fmt.Errorf(
				"insert final batch %d: %w",
				batchNumber,
				err,
			)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit result transaction: %w", err)
	}

	committed = true
	return nil
}

func insertQueryResultsBatch(
	tx *sql.Tx,
	results []QueryResults,
) error {
	if len(results) == 0 {
		return nil
	}

	var query strings.Builder

	query.WriteString(`
		INSERT INTO QueryResults
			(query_type, timeEnded, duration_ms)
		VALUES
	`)

	args := make([]any, 0, len(results)*3)

	for i, result := range results {
		if result.end.IsZero() {
			return fmt.Errorf(
				"result at index %d has zero end time",
				i,
			)
		}

		if result.duration < 0 {
			return fmt.Errorf(
				"result at index %d has negative duration: %s",
				i,
				result.duration,
			)
		}

		if i > 0 {
			query.WriteByte(',')
		}

		query.WriteString("(?, ?, ?)")

		durationMS := float64(result.duration) /
			float64(time.Millisecond)

		timeEnded := result.end.Format(
			"2006-01-02 15:04:05.000000",
		)

		args = append(
			args,
			result.qtype,
			timeEnded,
			durationMS,
		)
	}

	result, err := tx.Exec(query.String(), args...)
	if err != nil {
		return fmt.Errorf(
			"execute QueryResults insert for %d rows: %w",
			len(results),
			err,
		)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("read affected rows: %w", err)
	}

	if affected != int64(len(results)) {
		return fmt.Errorf(
			"unexpected affected rows: expected %d, got %d",
			len(results),
			affected,
		)
	}

	return nil
}
