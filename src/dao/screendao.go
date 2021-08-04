package dao

import (
	"model"
	"utils"
)

// AddScreen 增加Screen
func AddScreen(screen *model.Screen) error {
	sql := "insert into screen (name, rows, cols, img_path) values (?,?,?,?);"
	_, err := utils.DB.Exec(sql, screen.Name, screen.Rows, screen.Cols, screen.ImgPath)
	if err != nil {
		return err
	}
	sql2 := "select @@IDENTITY"
	err = utils.DB.QueryRow(sql2).Scan(&screen.ID)
	if err != nil {
		return err
	}
	//再生成座位
	err = AddSeatsByScreen(screen)
	if err != nil {
		return err
	}
	return err
}

// DeleteScreenByID 删除Screen
func DeleteScreenByID(screenID int) error {
	//先删除计划
	err := DeletePlansByScreenID(screenID)
	if err != nil {
		return err
	}
	//先删除座位
	err = DeleteSeatsByScreenID(screenID)
	if err != nil {
		return err
	}
	sql := "delete from screen where id = ?;"
	_, err = utils.DB.Exec(sql, screenID)
	if err != nil {
		return err
	}
	return err
}

// GetScreenByID 通过ID获取Screen信息
func GetScreenByID(screenID int) (*model.Screen, error) {
	sql := "select id, name, rows, cols, img_path from screen where id = ?;"
	var screen model.Screen
	err := utils.DB.QueryRow(sql, screenID).Scan(&screen.ID, &screen.Name, &screen.Rows, &screen.Cols, &screen.ImgPath)
	if err != nil {
		return nil, err
	}
	return &screen, err
}

// GetPageScreens 获取分页影厅
func GetPageScreens(pageNo int) (*model.Page, error) {
	page := model.Page{PageSize: PAGESIZE, PageNo: pageNo}
	//获取总记录数和总页数
	sql := "select count(*) from screen;"
	err := utils.DB.QueryRow(sql).Scan(&page.TotalRecord)
	if err != nil {
		return nil, err
	}
	page.TotalPageNo = page.TotalRecord / page.PageSize
	if page.TotalRecord%page.PageSize != 0 {
		page.TotalPageNo++
	}
	//通过limit获取影厅
	sql = "select id, name, rows, cols, img_path from screen limit ?,?;"
	rows, err := utils.DB.Query(sql, (page.PageNo-1)*page.PageSize, page.PageSize)
	if err != nil {
		return nil, err
	}
	var screens []*model.Screen
	for rows.Next() {
		var screen model.Screen
		err := rows.Scan(&screen.ID, &screen.Name, &screen.Rows, &screen.Cols, &screen.ImgPath)
		if err != nil {
			return nil, err
		}
		screens = append(screens, &screen)
	}
	page.Screens = screens
	return &page, err
}

// GetAllScreens 获取所有电影信息
func GetAllScreens() ([]*model.Screen, error) {
	sql := "select id, name, rows, cols, img_path  from screen;"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var screens []*model.Screen
	for rows.Next() {
		var screen model.Screen
		err := rows.Scan(&screen.ID, &screen.Name, &screen.Rows, &screen.Cols, &screen.ImgPath)
		if err != nil {
			return nil, err
		}
		screens = append(screens, &screen)
	}
	return screens, err
}
