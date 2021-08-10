package server

import (
	"dao"
)

// UpdateBox 更新票房
func UpdateBox(price float64, movieID int) error {
	box, err := dao.GetBoxByMovieID(movieID)
	if err != nil {
		return err
	}
	box.TotalAmount += price
	return dao.UpdateBox(box)
}
