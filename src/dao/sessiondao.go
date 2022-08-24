package dao

import (
	"cinema/model"
	"cinema/utils"
)

const (
	CookieKEY = "user"
)

// AddAndUpdateSession 更新session数据
func AddAndUpdateSession(session *model.Session) error {
	sql := "delete from session where user_id = ?;"
	_, err := utils.DB.Exec(sql, session.UserID)
	if err != nil {
		return err
	}
	sql = "insert into session(id, user_id) values(?,?) "
	_, err = utils.DB.Exec(sql, session.ID, session.UserID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession 通过sessionID删除session
func DeleteSession(sessionID string) error {
	sql := "delete from session where id = ?;"
	_, err := utils.DB.Exec(sql, sessionID)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionByID 通过sessionID获取session
func GetSessionByID(sessionID string) (*model.Session, error) {
	sqlStr := "select  id, user_id from session where id = ?"
	var session model.Session
	err := utils.DB.QueryRow(sqlStr, sessionID).Scan(&session.ID, &session.UserID)
	if err != nil {
		return nil, err
	}
	user, err := GetUserByID(session.UserID) // 填充其他字段
	if err != nil {
		return nil, err
	}
	session.Root = user.Root
	session.UserName = user.Name
	return &session, nil
}
