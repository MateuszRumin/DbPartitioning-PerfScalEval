package config

import (
	"database/sql"
	"time"
)

var (
	NumConnections int
	Logs           = false
	Prometeus      = false
	Action         = []string{"sdasdasd", "sdfsafsad"}
	Connections    []*sql.DB
	testTime       time.Duration
)
