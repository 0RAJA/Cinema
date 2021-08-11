package model

// Comment 评论
type Comment struct {
	ID       int    `json:"id,omitempty"`
	MovieID  int    `json:"movie_id,omitempty"`
	UserID   int    `json:"user_id,omitempty"`
	Cnt      int    `json:"cnt,omitempty"`
	Time     string `json:"time,omitempty"`
	Text     string `json:"text,omitempty"` //--
	UserName string `json:"user_name,omitempty"`
	ImgPath  string `json:"img_path,omitempty"`
}
