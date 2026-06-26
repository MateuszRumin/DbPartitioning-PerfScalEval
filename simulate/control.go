package main

import (
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

func multiThreadUpdate(workersCount int, deadline time.Time, idb int, idc int, idph int, idp int, idu int) {
	var wg sync.WaitGroup

	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			updateTest(id, deadline, idb, idc, idph, idp, idu)
		}(i)
	}

}

func multiThreadInsert(workersCount int, deadline time.Time) {
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			insertTest(deadline, id)
		}(i)
	}

}
