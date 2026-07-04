package main

import (
	"fmt"
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

	threads := 50
	var wg sync.WaitGroup
	start := time.Now()
	deadline := time.Now().Add(10 * time.Minute)

	for i := 0; i < threads; i++ {
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
	_, err = db.Exec("Insert INTO Tests (name,timeStart,timeEnd) values (?,?,?)", fmt.Sprintf("SR %d P ", threads), start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Printf("result insert error: %v", err)
	}

}
