package model

// Box 票房
type Box struct {
	MovieID     int     `json:"movie_id,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"` //--
	MovieName   string  `json:"movie_name,omitempty"`
	ImgPath     string  `json:"img_path,omitempty"`
}
