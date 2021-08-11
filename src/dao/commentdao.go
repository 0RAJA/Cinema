package dao

import (
	"model"
	"utils"
)

// AddComment 新增评论
func AddComment(comment *model.Comment) error {
	sql := "insert into comment (text, user_id, movie_id, cnt, time) values (?,?,?,?,?);"
	_, err := utils.DB.Exec(sql, comment.Text, comment.UserID, comment.MovieID, comment.Cnt, comment.Time)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCommitByID 删除评论
func DeleteCommitByID(commentID int) error {
	sql := "delete from comment where id = ?;"
	_, err := utils.DB.Exec(sql, commentID)
	if err != nil {
		return err
	}
	return nil
}

// GetCommentByID 获取指定评论
func GetCommentByID(commentID int) (*model.Comment, error) {
	sql := "select comment.id, text, user_id, movie_id, cnt, time,u.name,u.img_path from comment join users u on comment.user_id = u.id where comment.id = ?;"
	var comment model.Comment
	err := utils.DB.QueryRow(sql, commentID).Scan(&comment.ID, &comment.Text, &comment.UserID, &comment.MovieID, &comment.Cnt, &comment.Time, &comment.UserName, &comment.ImgPath)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetAllComments 获取全部评论
func GetAllComments() ([]*model.Comment, error) {
	sql := "select comment.id, text, user_id, movie_id, cnt, time,u.name,u.img_path from comment join users u on comment.user_id = u.id order by cnt desc ;"
	var comments []*model.Comment
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.ID, &comment.Text, &comment.UserID, &comment.MovieID, &comment.Cnt, &comment.Time, &comment.UserName, &comment.ImgPath)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, err
}

func UpdateComment(comment *model.Comment) error {
	sql := "update comment set text = ?,cnt = ? where id = ?"
	_, err := utils.DB.Exec(sql, comment.Text, comment.Cnt, comment.ID)
	if err != nil {
		return err
	}
	return nil
}
