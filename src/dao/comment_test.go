package dao

import (
	"fmt"
	"log"
	"testing"
	"time"

	"cinema/model"
)

func TestAddComment(t *testing.T) {
	comment := model.Comment{
		MovieID:  2,
		UserID:   13,
		Cnt:      0,
		Time:     time.Now().Format(model.PlanTimeFormat),
		Text:     "HHH",
		UserName: "Rajaa",
		ImgPath:  "www",
	}
	err := AddComment(&comment)
	if err != nil {
		return
	}
}

func TestDeleteCommitByID(t *testing.T) {
	err := DeleteCommitByID(1)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetAllComments(t *testing.T) {
	comments, err := GetAllComments()
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range comments {
		fmt.Println(v)
	}
}

func TestGetCommentByID(t *testing.T) {
	commentID := 2
	comment, err := GetCommentByID(commentID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(comment)
}

func TestUpdateComment(t *testing.T) {
	comment, _ := GetCommentByID(2)
	comment.Cnt = 2
	err := UpdateComment(comment)
	if err != nil {
		log.Println(err)
		return
	}
}
