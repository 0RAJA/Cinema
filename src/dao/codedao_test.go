package dao

import (
	"fmt"
	"model"
	"testing"
	"utils"
)

func TestAddCode(t *testing.T) {
	code := model.Code{
		Code: utils.RandStringRunes(),
		Root: model.RootAdmin,
	}
	err := AddCode(&code)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestCheckCode(t *testing.T) {
	code, err := CheckCode("juHOhtXeIh")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*code)
}

func TestDeleteCode(t *testing.T) {
	err := DeleteCode("juHOhtXeIh")
	if err != nil {
		fmt.Println(err)
		return
	}
}
