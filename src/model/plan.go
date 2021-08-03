package model

// Plan 演出计划
type Plan struct {
	ID         int
	ScreenID   int
	MovieID    int
	UpTime     string
	DownTime   string
	Price      float64 //--
	ScreenName string
	MovieName  string
}

// PlanCondition 演出计划查询条件
type PlanCondition struct {
	ScreenID string
	MovieID  string
	UpTime   string
}
