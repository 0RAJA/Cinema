package dao

import (
	"fmt"
	"testing"

	"cinema/model"
)

func TestAddScreen(t *testing.T) {
	screen := model.Screen{
		Name:    "1",
		Rows:    10,
		Cols:    10,
		ImgPath: "src/views/img/screen/1.png",
	}
	err := AddScreen(&screen)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDeleteScreenByID(t *testing.T) {
	err := DeleteScreenByID(5)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetScreenByID(t *testing.T) {
	screen, err := GetScreenByID(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(screen)
}

func TestGetPageScreens(t *testing.T) {
	page, err := GetPageScreens(1)
	if err != nil {
		return
	}
	fmt.Println(page)
	for i := 0; i < len(page.Screens); i++ {
		fmt.Println(page.Screens[i])
	}
}
