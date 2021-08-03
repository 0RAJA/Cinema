package model

// Box 票房
type Box struct {
	ID          int
	MovieID     int
	TotalAmount float64 //--
	MovieName   string
	ImgPath     string
}
