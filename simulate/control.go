package main

import (
	"database/sql"
	sqlgen "simulate/sqlgenerate"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type QueryResults struct {
	qtype    string
	end      time.Time
	duration time.Duration
}

func GetIDs(db *sql.DB, query string) ([]int, error) {
	rows, err := db.Query(query)
	if err != nil {
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func multiThreadSelect(workersCount int, deadline time.Time, idp int, idu int) {
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			r := newWorkerRand()
			wg := sqlgen.NewWorkerGenerator(r)
			selectTest(id, r, wg, idp, idu, deadline)
		}(i)
	}
	wg.Wait()
}

func multiThreadUpdate(workersCount int, deadline time.Time, idb []int, idc []int, idph []int, idp []int, idu []int) {
	var wg sync.WaitGroup

	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			updateTest(deadline, idb, idc, idph, idp, idu)
		}(i)
	}

}

func multiThreadInsert(workersCount int, deadline time.Time, idb []int, idc []int, idph []int, idp []int, idu []int) {
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			insertTest(deadline, idb, idc, idph, idp, idu)
		}(i)
	}

}
