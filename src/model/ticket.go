package model

const (
	TicketUnsold = iota + 2 //未售
	TicketLocked            //锁定
	TicketSold              //已售
	TicketError             //损坏
	TicketNone              //走廊
)

// Ticket 电影票
type Ticket struct {
	ID         int
	ScreenID   int
	MovieID    int
	PlanID     int
	Row        int
	Col        int
	State      int
	UserID     int //--
	ScreenName string
	MovieName  string
	UserName   string
	UpTime     string
	DownTime   string
	Price      float64
}
