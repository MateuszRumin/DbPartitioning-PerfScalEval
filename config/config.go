package config

import (
	"database/sql"
	"time"
)

var (
	NumConnections   int  = 1
	Logs             bool = false
	Prometeus        bool = false
	Action                = []string{"none", "Select", "Create", "Mix"}
	Connections      []*sql.DB
	TestTime         time.Duration
	TestType         string = "none"
	TestName         string = "none"
	ConectionWorking bool   = false
	Timeout          time.Duration
)
