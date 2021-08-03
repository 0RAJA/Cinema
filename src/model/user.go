package model

const (
	RootClient  = iota //顾客
	RootManager        //经理
	RootAdmin          //管理员
)

// User 用户信息
type User struct {
	ID       int
	Name     string
	Password string
	Email    string
	Root     int
	ImgPath  string
}
