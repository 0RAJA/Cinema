package model

const MovieTimeFormat = "2006-1-2"

// Movie 电影信息
type Movie struct {
	ID       int     `json:"movieID,omitempty"`
	Name     string  `json:"name,omitempty"`
	UpDate   string  `json:"upDate,omitempty"` //时间
	Director string  `json:"director,omitempty"`
	Type     string  `json:"type,omitempty"`
	Area     string  `json:"area,omitempty"`
	Star     string  `json:"star,omitempty"`
	Length   int     `json:"length,omitempty"`
	Intro    string  `json:"intro,omitempty"`
	ImgPath  string  `json:"imgPath,omitempty"`
	Score    float64 `json:"score,omitempty"`
}
