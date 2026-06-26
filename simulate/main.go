package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	workersCountSelect := 7
	workersCountInsert := 4
	workersCountUpdate := 2

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

	deadline := time.Now().Add(1 * time.Hour)
	wg.Add(3)

	start := time.Now()

	go func() {
		defer wg.Done()
		multiThreadSelect(workersCountSelect, deadline, idp, idu)
	}()

	go func() {
		defer wg.Done()
		multiThreadInsert(workersCountInsert, deadline)

	}()

	go func() {
		defer wg.Done()
		multiThreadUpdate(workersCountUpdate, deadline, idb, idc, idph, idp, idu)
	}()

	wg.Wait()

	stop := time.Now()

	db2, err := slc()
	if err != nil {

		return
	}
	defer db2.Close()
	db2.Query(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('%s','%s','%s')", "Simulate 1h 742 np", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

}
