package main

import (
	"database/sql"
	"fmt"
	"strings"
)

func saveQueryResults(
	db *sql.DB, resultsByWorker [][]QueryResults,
) error {

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf(
			"begin result transaction: %w",
			err,
		)
	}

	committed := false

	defer func() {
		if !committed {
			_ = tx.Rollback()
		}
	}()

	const batchSize = 1000

	batch := make([]QueryResults, 0, batchSize)

	for _, workerResults := range resultsByWorker {
		for _, result := range workerResults {
			batch = append(batch, result)

			if len(batch) == batchSize {
				if err := insertQueryResultsBatch(
					tx,
					batch,
				); err != nil {
					return err
				}

				batch = batch[:0]
			}
		}
	}

	if len(batch) > 0 {
		if err := insertQueryResultsBatch(
			tx,
			batch,
		); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf(
			"commit result transaction: %w",
			err,
		)
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

	query.WriteString(
		`INSERT INTO QueryResults
		(query_type, timeEnded, duration_ms)
		VALUES `,
	)

	args := make([]any, 0, len(results)*3)

	for i, result := range results {
		if i > 0 {
			query.WriteByte(',')
		}

		query.WriteString("(?,?,?)")

		durationMS := float64(
			result.duration.Microseconds(),
		) / 1000.0

		args = append(
			args,
			result.qtype,
			result.end,
			durationMS,
		)
	}

	_, err := tx.Exec(query.String(), args...)
	if err != nil {
		return fmt.Errorf(
			"insert result batch: %w",
			err,
		)
	}

	return nil
}
