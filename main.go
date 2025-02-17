package main

import (
	"bufio"
	"fmt"
	"os"

	"perfscaleval/config"
	"perfscaleval/dbconects"
	"perfscaleval/test"
)

func main() {
	fmt.Println("Start Main")

	config.NumConnections = 1
	dbconects.SetConnections()

	test.OneConAllTable()

	fmt.Println("Finish Main")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	//resourceusage.LogUsage()
}
