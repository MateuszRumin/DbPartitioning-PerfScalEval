package main

import "fmt"

func main() {

	db, err := setConnection()
	if err != nil {
		fmt.Printf("Error setting connection: %v\n", err)
		return
	}
	defer db.Close()

	multiThreadConnection(db)
}
