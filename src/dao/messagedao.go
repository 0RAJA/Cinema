package dao

import (
	"model"
	"utils"
)

// AddAndUpdateMessage 增加并更新验证码
func AddAndUpdateMessage(message *model.Message) error {
	defer rwmutexMessage.Unlock()
	rwmutexMessage.Lock()
	sql := "delete from message where email = ?;"
	_, err := utils.DB.Exec(sql, message.Email)
	if err != nil {
		return err
	}
	sql = "insert into message VALUES (?,?)"
	_, err = utils.DB.Exec(sql, message.Email, message.Str)
	if err != nil {
		return err
	}
	return nil
}

// CheckMessage 检查验证码是否存在
func CheckMessage(message *model.Message) (*model.Message, error) {
	defer rwmutexMessage.RUnlock()
	rwmutexMessage.RLock()
	sql := "select email, str from message where email = ? and str = ?;"
	var rM model.Message
	err := utils.DB.QueryRow(sql, message.Email, message.Str).Scan(&rM.Email, &rM.Str)
	if err != nil {
		return nil, err
	}
	return &rM, nil
}

// DeleteMessage 删除邀请码
func DeleteMessage(email string) error {
	defer rwmutexMessage.Unlock()
	rwmutexMessage.Lock()
	sql := "delete from message where email = ?"
	_, err := utils.DB.Exec(sql, email)
	if err != nil {
		return err
	}
	return nil
}
