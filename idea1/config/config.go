package config

import (
	"database/sql"
	"flag"
	"time"
)

var (
	NumConnections   int       = 1
	Logs             bool      = false
	Prometeus        bool      = false
	Action                     = []string{"none", "Select", "Create", "Mix"}
	Connections      []*sql.DB = nil
	TestTime         time.Duration
	TestType         string = "none"
	TestName         string = "none"
	ConectionWorking bool   = false
	Timeout          time.Duration
	BuforCreateSize  int = 20
)
var Interval = flag.Int("interval", 1, "Interval in seconds between actions")
