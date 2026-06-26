package main

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func checkConnectionAndRunTest(id int) {
	// if id > 10 && id < 21 {
	// 	time.Sleep(20 * time.Second)
	// }
	db, err := setConnection()
	if err != nil {
		log.Fatalf("Nie udało się połączyć z bazą danych: %v", err)
	} else {
		fmt.Println("Połączenie z bazą danych działa poprawnie. Wątek: ", id)
		for l := 0; l < 10; l++ {

			testDb(db, id)
		}
	}
	defer db.Close()

}

func multiThread() {

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			checkConnectionAndRunTest(id)
		}(i)
	}

	wg.Wait()

}
