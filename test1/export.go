package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func exportCSV(path string, results []QueryResult) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{
		"query",
		"duration_ns",
		"first_row_ns",
		"rows",
		"explain_rows",
		"partitions",
	})

	for _, r := range results {
		w.Write([]string{
			r.Query,
			strconv.FormatInt(r.Duration.Nanoseconds(), 10),
			strconv.FormatInt(r.FirstRow.Nanoseconds(), 10),
			strconv.Itoa(r.Rows),
			strconv.Itoa(r.ExplainRows),
			r.Partitions,
		})
	}

	return nil
}
