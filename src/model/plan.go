package model

const (
	// PlanTimeFormat 计划时间
	PlanTimeFormat   = `2006-01-02 15:04:05`
	PlanTimeFormatOK = `2006-01-02T15:04:05`
)

// Plan 演出计划
type Plan struct {
	ID             int     `json:"id,omitempty"`
	ScreenID       int     `json:"screen_id,omitempty"`
	MovieID        int     `json:"movie_id,omitempty"`
	UpTime         string  `json:"up_time,omitempty"`
	DownTime       string  `json:"down_time,omitempty"`
	Price          float64 `json:"price,omitempty"` //--
	ScreenName     string  `json:"screen_name,omitempty"`
	MovieName      string  `json:"movie_name,omitempty"`
	UpTimeFormat   string  `json:"up_time_format,omitempty"` //给前端看的
	DownTimeFormat string  `json:"down_time_format,omitempty"`
	IsTimeOut      bool    `json:"is_time_out,omitempty"`
}

// MovieAndScreen 当前存在的电影和影厅
type MovieAndScreen struct {
	Movies  []*Movie  `json:"movies,omitempty"`
	Screens []*Screen `json:"screens,omitempty"`
}

// PlanTimeMessage 前端传过来的时间校验信息
type PlanTimeMessage struct {
	MovieID      int    `json:"movie_id,omitempty"`
	ScreenID     int    `json:"screen_id,omitempty"`
	UpTimeFormat string `json:"up_time_format,omitempty"`
}

// ReplyTimeMessage 回复时间校验信息
type ReplyTimeMessage struct {
	DownTimeFormat string `json:"down_time_format,omitempty"`
	IsOK           bool   `json:"is_ok,omitempty"`
}

// AddOrUpdatePlanMessage 增加或者修改plan时前端传输的数据
type AddOrUpdatePlanMessage struct {
	ID             int     `json:"id,omitempty"`
	ScreenID       int     `json:"screen_id,omitempty"`
	MovieID        int     `json:"movie_id,omitempty"`
	UpTimeFormat   string  `json:"up_time_format,omitempty"`
	DownTimeFormat string  `json:"down_time_format,omitempty"`
	Price          float64 `json:"price,omitempty"` //--
}
