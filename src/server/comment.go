package server

import (
	"dao"
	"model"
)

func AddComment(comment *model.Comment) error {
	return dao.AddComment(comment)
}

func DeleteComment(commentID int) error {
	return dao.DeleteCommitByID(commentID)
}

func GetCommentByID(commentID int) (*model.Comment, error) {
	return dao.GetCommentByID(commentID)
}

func GetAllComments(movieID int) ([]*model.Comment, error) {
	return dao.GetAllComments(movieID)
}

func IsSelfComment(commentID, userID int) (bool, error) {
	comment, err := dao.GetCommentByID(commentID)
	if err != nil {
		return false, err
	}
	return comment.UserID == userID, nil
}

func UpdateComment(comment *model.Comment) error {
	return dao.UpdateComment(comment)
}
