package main

import (
	"fmt"
	"log"
	sqlgen "simulate/sqlgenerate"
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

	idb, errb := GetIDs(db, "SELECT id FROM badges")
	idc, errc := GetIDs(db, "SELECT id FROM comments")
	idph, errph := GetIDs(db, "SELECT id FROM post_history")
	idp, errp := GetIDs(db, "SELECT id FROM posts")
	idu, erru := GetIDs(db, "SELECT id FROM users")
	if errb != nil || errc != nil || errph != nil || errp != nil || erru != nil {
		fmt.Printf("Brak danych o indeksach")
		return
	}

	idpi, idui, err := sqlgen.GenerateInitValues()
	if err != nil {
		log.Println("Błąd generacji zapytania")
		return
	}

	deadline := time.Now().Add(30 * time.Minute)
	wg.Add(3)

	start := time.Now()

	go func() {
		defer wg.Done()
		multiThreadSelect(workersCountSelect, deadline, idpi, idui)
	}()

	go func() {
		defer wg.Done()
		multiThreadInsert(workersCountInsert, deadline, idb, idc, idph, idp, idu)

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
	db2.Exec(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('%s','%s','%s')",
		"Simulate 30m 742 np", start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

}
