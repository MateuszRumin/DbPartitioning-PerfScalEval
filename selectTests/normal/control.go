package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	sqlgen "normal/sqlgenerate"

	_ "github.com/go-sql-driver/mysql"
)

func multiThreadConnection() {
	idp, idu, err := sqlgen.GenerateInitValues()
	if err != nil {
		log.Println("Błąd generacji zapytania")
		return
	}

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			r := newWorkerRand()
			wg := sqlgen.NewWorkerGenerator(r)
			wantConnection(id, r, wg, idp, idu)
		}(i)
	}

	wg.Wait()

	stop := time.Now()

	db, err := slc()
	if err != nil {

		return
	}
	defer db.Close()
	db.Query(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('%s','%s','%s')", "Select normal workload 1h 10threads", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

}
