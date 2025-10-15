package confmodel

type TestPlan struct {
	PlanName string      `json:"plan_name"`
	Comments string      `json:"comments"`
	Group    []TestGroup `json:"Groups"`
}

type TestGroup struct {
	Name         string     `json:"name"`
	Comments     string     `json:"comments"`
	Action       string     `json:"action"`
	ThreadNumber int        `json:"threads"`
	RampUp       int        `json:"ramp_up"`
	SameUserIter bool       `json:"same_user_on_iter"`
	Loops        int        `json:"loops"`
	Infinite     bool       `json:"infinite"`
	LifeTime     int        `json:"lifetime"`
	Duration     int        `json:"duration"`
	Steps        []TestStep `json:"Steps"`
}

type TestStep struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

var CreateMode bool = true
var GroupExist bool = false
var ChooseLayOut int = 0
var CurrentPlan *TestPlan
var CurrentGroup *TestGroup
var CurrentStep *TestStep
var Plan []TestPlan
