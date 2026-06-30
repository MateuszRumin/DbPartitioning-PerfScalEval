package main

import (
	"log"
	"sync"
	"time"

	sqlgen "range/sqlgenerate"

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
	deadline := time.Now().Add(1 * time.Hour)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			r := newWorkerRand()
			wg := sqlgen.NewWorkerGenerator(r)
			wantConnection(deadline, id, r, wg, idp, idu)
		}(i)
	}

	wg.Wait()

	stop := time.Now()

	db, err := slc()
	if err != nil {

		return
	}
	defer db.Close()
	_, err = db.Exec("Insert INTO Tests (name,timeStart,timeEnd) values ('?','?','?')", "Select range workload 1h 10threads p", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Printf("result insert error: %v", err)
	}

}
