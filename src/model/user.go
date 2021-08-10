package model

const (
	RootClient  = iota //顾客
	RootManager        //经理
	RootAdmin          //管理员
)

// User 用户信息
type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Root     int    `json:"root,omitempty"`
	ImgPath  string `json:"imgPath,omitempty"`
}
