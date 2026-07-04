package main

type ConnectionSetting struct {
	Connection   int
	Enable       bool
	ModelInfo    string
	User         string
	Host         string
	Port         string
	Password     string
	DatabaseName string
}

var Connection = ConnectionSetting{
	Connection:   1,
	Enable:       false,
	ModelInfo:    "",
	User:         "",
	Host:         "",
	Port:         "",
	Password:     "",
	DatabaseName: "",
}
