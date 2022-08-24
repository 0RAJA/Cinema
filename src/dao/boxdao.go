package dao

import (
	"cinema/model"
	"cinema/utils"
)

// UpdateBox 更新票房
func UpdateBox(box *model.Box) error {
	defer rwMutexBox.RUnlock()
	rwMutexBox.RLock()
	sql := "update box set total_amount = ? where movie_id = ?"
	_, err := utils.DB.Exec(sql, box.TotalAmount, box.MovieID)
	if err != nil {
		return err
	}
	return err
}

// AddBox 新建票房
func AddBox(movieID int) error {
	defer rwMutexBox.Unlock()
	rwMutexBox.Lock()
	sql := "insert into box (movie_id,total_amount) values(?,?)"
	_, err := utils.DB.Exec(sql, movieID, 0)
	if err != nil {
		return err
	}
	return err
}

// GetBoxByMovieID 通过movieID获取票房信息
func GetBoxByMovieID(movieID int) (*model.Box, error) {
	defer rwMutexBox.RUnlock()
	rwMutexBox.RLock()
	sql := "select box.movie_id,box.total_amount,movie.name,movie.img_path from box join movie on box.movie_id = movie.id where box.movie_id = ?;"
	var box model.Box
	err := utils.DB.QueryRow(sql, movieID).Scan(&box.MovieID, &box.TotalAmount, &box.MovieName, &box.ImgPath)
	if err != nil {
		return nil, err
	}
	return &box, err
}

// GetAllBoxes 获取所有票房
func GetAllBoxes() ([]*model.Box, error) {
	defer rwMutexBox.RUnlock()
	rwMutexBox.RLock()
	sql := "select b.movie_id,b.total_amount,m.name,m.img_path from box b join movie m on m.id = b.movie_id"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var boxes []*model.Box
	for rows.Next() {
		var box model.Box
		err = rows.Scan(&box.MovieID, &box.TotalAmount, &box.MovieName, &box.ImgPath)
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, &box)
	}
	return boxes, err
}
