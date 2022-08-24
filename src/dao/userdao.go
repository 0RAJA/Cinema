package dao

import (
	"cinema/model"
	"cinema/utils"
)

// CheckUserNamePwd 检查用户名和密码
func CheckUserNamePwd(name, password string) (*model.User, error) {
	defer rwmutexUser.RUnlock()
	rwmutexUser.RLock()
	newPassword := password
	user := model.User{}
	sql := "select * from users where binary name = ? and password = ?;"
	err := utils.DB.QueryRow(sql, name, newPassword).Scan(&user.ID, &user.Name, &user.Password, &user.Root, &user.ImgPath, &user.Email)
	return &user, err
}

// IsNameOK 检查用户名是否可用
func IsNameOK(name string) (*model.User, error) {
	defer rwmutexUser.RUnlock()
	rwmutexUser.RLock()
	sql := "select id from users where binary name = ?;"
	user := model.User{}
	err := utils.DB.QueryRow(sql, name).Scan(&user.ID)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

// IsEmailOK 检查邮箱是否可用
func IsEmailOK(email string) (*model.User, error) {
	defer rwmutexUser.RUnlock()
	rwmutexUser.RLock()
	sql := "select id from users where binary email = ?;"
	user := model.User{}
	err := utils.DB.QueryRow(sql, email).Scan(&user.ID)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

// InsertUser 插入用户信息
func InsertUser(user *model.User) error {
	defer rwmutexUser.Unlock()
	rwmutexUser.Lock()
	newPassword := user.Password
	sql := "insert into users(name, password, root, img_path, email) VALUES(?,?,?,?,?)"
	_, err := utils.DB.Exec(sql, user.Name, newPassword, user.Root, user.ImgPath, user.Email)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *model.User) error {
	defer rwmutexUser.Unlock()
	rwmutexUser.Lock()
	newPassword := user.Password
	sql := "update users set name = ?,password = ? where id = ?"
	_, err := utils.DB.Exec(sql, user.Name, newPassword, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID 通过userID获取user信息
func GetUserByID(userID int) (*model.User, error) {
	defer rwmutexUser.RUnlock()
	rwmutexUser.RLock()
	user := model.User{}
	sql := "select id, name, password, root, img_path, email from users where id = ?;"
	err := utils.DB.QueryRow(sql, userID).Scan(&user.ID, &user.Name, &user.Password, &user.Root, &user.ImgPath, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail 通过邮箱获取邮箱信息
func GetUserByEmail(email string) (*model.User, error) {
	defer rwmutexUser.RUnlock()
	rwmutexUser.RLock()
	user := model.User{}
	sql := "select * from users where email = ?;"
	err := utils.DB.QueryRow(sql, email).Scan(&user.ID, &user.Name, &user.Password, &user.Root, &user.ImgPath, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateImg(user *model.User) error {
	defer rwmutexUser.Unlock()
	rwmutexUser.Lock()
	sql := "update users set img_path = ? where id = ?"
	_, err := utils.DB.Exec(sql, user.ImgPath, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmail(user *model.User) error {
	defer rwmutexUser.Unlock()
	rwmutexUser.Lock()
	sql := "update users set email = ? where id = ?"
	_, err := utils.DB.Exec(sql, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}
