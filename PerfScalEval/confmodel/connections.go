package confmodel

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

var Connection1 = ConnectionSetting{
	Connection:   1,
	Enable:       false,
	ModelInfo:    "",
	User:         "",
	Host:         "",
	Port:         "",
	Password:     "",
	DatabaseName: "",
}

var Connection2 = ConnectionSetting{
	Connection:   2,
	Enable:       false,
	ModelInfo:    "",
	User:         "",
	Host:         "",
	Port:         "",
	Password:     "",
	DatabaseName: "",
}
