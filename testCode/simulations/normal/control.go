package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	sqlgen "normal/sqlgenerate"

	_ "github.com/go-sql-driver/mysql"
)

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

func multiThreadConnection() {

	db, err := setConnection()
	if err != nil {
		log.Printf("Błąd połączenia: %v", err)
		return
	}
	defer db.Close()

	idp, errp := GetIDs(db, "SELECT id FROM posts")
	idu, erru := GetIDs(db, "SELECT id FROM users")
	if errp != nil || erru != nil {
		fmt.Printf("Brak danych o indeksach")
		return
	}
	threads := 50

	var wg sync.WaitGroup

	resultsByWorker := make([][]QueryResults, threads)

	start := time.Now()
	deadline := time.Now().Add(10 * time.Minute)

	for i := 0; i < threads; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			r := newWorkerRand()
			wg, err := sqlgen.NewWorkerGenerator(sqlgen.ReadGenerators)
			if err != nil {
				log.Printf("Błąd tworzenia generatora: %v", err)
				return
			}
			resultsByWorker[id] = wantConnection(deadline, id, r, wg, idp, idu)

		}(i)
	}

	wg.Wait()

	stop := time.Now()

	db2, err := slc()
	if err != nil {
		return
	}
	defer db2.Close()

	db2.Exec(fmt.Sprintf("Insert INTO Tests (name,timeStart,timeEnd) values ('SN %d Partycjonowana','%s','%s')",
		threads, start.Format("2006-01-02 15:04:05"), stop.Format("2006-01-02 15:04:05")))

	if err := saveQueryResults(db2, resultsByWorker); err != nil {
		log.Printf("result save error: %v", err)
	}

}
