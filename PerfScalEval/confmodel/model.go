package confmodel

type TestPlan struct {
	PlanName string      `json:"plan_name"`
	Comments string      `json:"comments"`
	Group    []TestGroup `json:"Groups"`
}

type TestGroup struct {
	IDGroup      int        `json:"id_group"`
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
	Order      int    `json:"order"`
	Name       string `json:"name"`
	QueryType  string `json:"query_type"`
	Query      string `json:"query"`
	SaveResult bool   `json:"save_result"`
}

var ProjectExist bool = false
var CurrentPlan *TestPlan
var CurrentGroup *TestGroup
var CurrentStep *TestStep
var MainPlanContentControl int = 0
var Plan []TestPlan
