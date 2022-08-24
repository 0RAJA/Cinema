package dao

import (
	"fmt"
	"testing"

	"cinema/model"
)

func TestInsertUser(t *testing.T) {
	user := model.User{
		Name:     "Raja",
		Password: "123456",
		Email:    "1647193241@qq.com",
		Root:     model.RootAdmin,
		ImgPath:  "/img/Raja.png",
	}
	err := InsertUser(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestCheckUserNamePwd(t *testing.T) {
	name, password := "Raja", "123456"
	user, err := CheckUserNamePwd(name, password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*user)
}

func TestIsNameOK(t *testing.T) {
	name1, name2 := "raja", "Raja"
	user, err := IsNameOK(name1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user)
	user2, err := IsNameOK(name2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user2)
}

func TestIsEmailOK(t *testing.T) {
	email := "1647193241@qq.com"
	user, err := IsEmailOK(email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user)
}

func TestUpdateUser(t *testing.T) {
	user := model.User{
		ID:       1,
		Name:     "raja",
		Password: "123456",
		Email:    "1647193241@qq.com",
		Root:     model.RootAdmin,
		ImgPath:  "/img/Raja.png",
	}
	err := UpdateUser(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
}
