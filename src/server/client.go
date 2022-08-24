package server

import (
	"cinema/dao"
	"cinema/model"
)

func GetAllBoxes() ([]*model.Box, error) {
	return dao.GetAllBoxes()
}
