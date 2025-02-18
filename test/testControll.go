package test

import "perfscaleval/config"

func TestControll() {

	if config.TestType == "Select" {
		selectAnalizeControl()
	}

}
