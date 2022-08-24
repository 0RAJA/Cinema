package server

import (
	"cinema/dao"
	"cinema/model"
)

// GetSeatsByScreenID 获取screenID
func GetSeatsByScreenID(screenID int) (*model.ScreenSeats, error) {
	seats, err := dao.GetSeatsByScreenID(screenID)
	if err != nil {
		return nil, err
	}
	screen, err := dao.GetScreenByID(screenID)
	if err != nil {
		return nil, err
	}
	var screenSeats = model.ScreenSeats{
		Screen: screen,
		Seats:  seats,
	}
	return &screenSeats, nil
}

// UpdateSeats 更新座位信息
func UpdateSeats(seatStates []*model.SeatState) error {
	for i := 0; i < len(seatStates); i++ {
		err := dao.UpdateSeat(seatStates[i])
		if err != nil {
			return err
		}
	}
	return nil
}
