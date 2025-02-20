package test

import (
	"fmt"
	"perfscaleval/config"
	"perfscaleval/dbconects"
)

func TestControll() {

	if config.NumConnections <= 0 {
		fmt.Println("Błąd flagi: nie można przeprowadzić testów bez połączenia z bazą danych")
		return
	}

	dbconects.SetConnections()
	if !config.ConectionWorking {
		fmt.Println("Brak połączenia")
		return
	}

	if config.TestType == "Select" {
		selectControl()
	} else if config.TestType == "Create" {
		createControl()
	}

}
