package dao

import (
	"model"
	"utils"
)

// AddSeatsByScreen 通过Screen增加座位
func AddSeatsByScreen(screen *model.Screen) error {
	defer rwmutexSeat.Unlock()
	rwmutexSeat.Lock()
	for i := 1; i <= screen.Rows; i++ {
		for j := 1; j <= screen.Cols; j++ {
			seat := model.Seat{
				ScreenID: screen.ID,
				Row:      i,
				Col:      j,
				State:    model.SeatOK,
			}
			err := AddSeat(&seat)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AddSeat 增加新座位
func AddSeat(seat *model.Seat) error {
	sql := "insert into seat (row, col, state, screen_id) values (?,?,?,?);"
	_, err := utils.DB.Exec(sql, seat.Row, seat.Col, seat.State, seat.ScreenID)
	if err != nil {
		return err
	}
	return err
}

// DeleteSeatsByScreenID 根据ScreenID删除相关Seats
func DeleteSeatsByScreenID(screenID int) error {
	defer rwmutexSeat.Unlock()
	rwmutexSeat.Lock()
	sql := "delete from seat where screen_id = ?;"
	_, err := utils.DB.Exec(sql, screenID)
	if err != nil {
		return err
	}
	return err
}

// GetSeatsByScreenID 通过ScreenID获取所有Seat
func GetSeatsByScreenID(screenID int) ([]*model.Seat, error) {
	defer rwmutexSeat.RUnlock()
	rwmutexSeat.RLock()
	sql := "select id, row, col, state, screen_id from seat where screen_id = ?;"
	rows, err := utils.DB.Query(sql, screenID)
	if err != nil {
		return nil, err
	}
	var seats []*model.Seat
	for rows.Next() {
		var seat model.Seat
		err = rows.Scan(&seat.ID, &seat.Row, &seat.Col, &seat.State, &seat.ScreenID)
		if err != nil {
			return nil, err
		}
		seats = append(seats, &seat)
	}
	return seats, err
}

// UpdateSeat 更新座位信息
func UpdateSeat(seatState *model.SeatState) error {
	defer rwmutexSeat.Unlock()
	rwmutexSeat.Lock()
	sql := "update seat set state = ? where id = ?"
	_, err := utils.DB.Exec(sql, seatState.State, seatState.ID)
	if err != nil {
		return err
	}
	return err
}
