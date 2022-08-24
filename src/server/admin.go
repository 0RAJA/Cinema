package server

import (
	"cinema/dao"
	"cinema/model"
	"cinema/utils"
)

func GetAllCodes() ([]*model.Code, error) {
	return dao.GetAllCodes()
}

func DeleteCodeByID(code string) error {
	return dao.DeleteCode(code)
}

func AddCode(root int) (string, error) {
	code := model.Code{
		Code: utils.RandStringRunes(),
		Root: root,
	}
	err := dao.AddCode(&code)
	if err != nil {
		return "", err
	}
	return code.Code, nil
}
