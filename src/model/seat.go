package model

const (
	SeatOK    = iota + 2 //可以生成票的座位
	SeatError            //不能生成票的坏座位
	SeatNone             //走廊
)

// Seat 座位
type Seat struct {
	ID       int `json:"id,omitempty"`
	ScreenID int `json:"screenID,omitempty"`
	Row      int `json:"row,omitempty"`
	Col      int `json:"col,omitempty"`
	State    int `json:"state,omitempty"`
}

// ScreenSeats 传给前端的结构
type ScreenSeats struct {
	Screen *Screen `json:"screen"`
	Seats  []*Seat `json:"seats,omitempty"`
}

// SeatState 后台给前端传的结构
type SeatState struct {
	ID    int `json:"id,omitempty"`
	State int `json:"state,omitempty"`
}
