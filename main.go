package main

import (
	"bufio"
	"fmt"
	"os"

	"perfscaleval/config"
	//"perfscaleval/dbconects"
	//"perfscaleval/test"
	"perfscaleval/test"
)

func main() {
	fmt.Println("Start Main")

	// config.NumConnections = 1
	// dbconects.SetConnections()
	config.NumConnections = 10
	// test.OneConAllTable()
	config.TestType = "Select"
	config.TestName = "idselcectPostLimit"

	test.TestControll()
	fmt.Println("Finish Main")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	//resourceusage.LogUsage()
}
