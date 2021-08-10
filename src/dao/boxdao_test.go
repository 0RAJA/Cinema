package dao

import (
	"fmt"
	"model"
	"testing"
)

func TestAddBox(t *testing.T) {
	err := AddBox(2)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestUpdateBox(t *testing.T) {
	err := UpdateBox(&model.Box{
		MovieID:     2,
		TotalAmount: 444,
		MovieName:   "",
		ImgPath:     "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetBoxByMovieID(t *testing.T) {
	box, err := GetBoxByMovieID(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(box)
}
