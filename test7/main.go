package main

import "sync"

func main() {
	var wg sync.WaitGroup

	workersCountSelect := 30
	workersCountInsert := 15
	workersCountUpdate := 10

	wg.Add(3)

	go func() {
		defer wg.Done()
		multiThreadSelect(workersCountSelect)
	}()

	go func() {
		defer wg.Done()
		multiThreadInsert(workersCountInsert)

	}()

	go func() {
		defer wg.Done()
		multiThreadUpdate(workersCountUpdate)
	}()

	wg.Wait()

}

//10 min test dodawanie
