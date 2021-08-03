package server

import (
	"dao"
	"model"
)

// GetPageScreens 获取分页影厅信息
func GetPageScreens(pageNo int) (*model.Page, error) {
	page, err := dao.GetPageScreens(1)
	if err != nil {
		return nil, err
	}
	if pageNo > page.TotalPageNo {
		pageNo = page.TotalPageNo
	}
	page, err = dao.GetPageScreens(pageNo)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// AddScreen 新增影厅
func AddScreen(screen *model.Screen) error {
	return dao.AddScreen(screen)
}

// DeleteScreenByID 删除影厅
func DeleteScreenByID(screenID int) error {
	return dao.DeleteScreenByID(screenID)
}
