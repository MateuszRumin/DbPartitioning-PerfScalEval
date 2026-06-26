package main

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func multiThread() {

	var wg sync.WaitGroup
	deadline := time.Now().Add(10 * time.Minute)

	start := time.Now()

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			checkConnectionAndRunTest(deadline, id)
		}(i)
	}

	wg.Wait()

	stop := time.Now()

	db2, err := slc()
	if err != nil {
		return
	}
	defer db2.Close()

	db2.Query(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('%s','%s','%s')",
		"Insert random 30min 10 threads np", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

}
