package main

import (
	"bufio"
	"fmt"
	"os"

	"perfscaleval/config"
	"perfscaleval/dbconects"
)

func main() {
	fmt.Println("placeholder")

	config.NumConnections = 1
	dbconects.SetConnections()
	fmt.Println("placeholder")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	//resourceusage.LogUsage()
}
