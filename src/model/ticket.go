package model

const (
	TicketUnsold = 2 //未售
	TicketSold   = 3 //已售
	TicketError  = 4 //损坏
	TicketNone   = 5 //走廊
)

// Ticket 电影票
type Ticket struct {
	ID         int     `json:"id,omitempty"`
	ScreenID   int     `json:"screen_id,omitempty"`
	MovieID    int     `json:"movie_id,omitempty"`
	PlanID     int     `json:"plan_id,omitempty"`
	Row        int     `json:"row,omitempty"`
	Col        int     `json:"col,omitempty"`
	State      int     `json:"state,omitempty"`
	UserID     int     `json:"user_id,omitempty"` //--
	ScreenName string  `json:"screen_name,omitempty"`
	MovieName  string  `json:"movie_name,omitempty"`
	UserName   string  `json:"user_name,omitempty"`
	UpTime     string  `json:"up_time,omitempty"`
	DownTime   string  `json:"down_time,omitempty"`
	Price      float64 `json:"price,omitempty"`
	IsTimeOut  bool    `json:"is_time_out,omitempty"`
}
