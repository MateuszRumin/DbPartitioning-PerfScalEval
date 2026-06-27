package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func multiThread() {

	var wg sync.WaitGroup
	deadline := time.Now().Add(10 * time.Minute)

	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	var idb, idc, idph, idp, idu int
	errb := db.QueryRow("SELECT COUNT(*) FROM badges").Scan(&idb)
	errc := db.QueryRow("SELECT COUNT(*) FROM comments").Scan(&idc)
	errph := db.QueryRow("SELECT COUNT(*) FROM post_history").Scan(&idph)
	errp := db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&idp)
	erru := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&idu)
	if errb != nil || errc != nil || errph != nil || errp != nil || erru != nil {
		fmt.Printf("Brak danych o indeksach")
		return
	}

	start := time.Now()

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			checkConnectionAndRunTest(deadline, id, idb, idc, idph, idp, idu)
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
