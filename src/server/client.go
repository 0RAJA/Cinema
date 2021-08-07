package server

import (
	"dao"
	"model"
)

func GetAllBoxes() ([]*model.Box, error) {
	return dao.GetAllBoxes()
}
