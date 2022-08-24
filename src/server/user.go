package server

import (
	"log"
	"net/http"

	"cinema/dao"
	"cinema/model"
	"github.com/gofrs/uuid"
)

const (
	CookieMaxAge = 60 * 100
)

// DeleteCookie 删除cookie
func DeleteCookie(cookieValue string) error {
	err := dao.DeleteSession(cookieValue) // 删除session
	if err != nil {
		return err
	}
	return err
}

func IsLogin(r *http.Request) (*model.Session, bool) {
	session := &model.Session{}
	var err error
	// 判断cookie
	cookie, _ := r.Cookie(dao.CookieKEY)
	if cookie != nil {
		// 获取cookie值
		cookieValue := cookie.Value
		// 去数据库查相关session
		session, err = dao.GetSessionByID(cookieValue)
		return session, err == nil
	}
	return session, false
}

// CheckUserNamePwd 验证用户名和密码
func CheckUserNamePwd(userName, password string) (*model.User, error) {
	return dao.CheckUserNamePwd(userName, password)
}

// CreatSession 创建session返回cookie
func CreatSession(user *model.User) (*http.Cookie, error) {
	// 创建session
	sessionID, _ := uuid.NewV4()
	session := model.Session{
		ID:       sessionID.String(),
		UserID:   user.ID,
		UserName: user.Name,
		Root:     user.Root,
	}
	// 添加session到服务器
	err := dao.AddAndUpdateSession(&session)
	if err != nil {
		return nil, err
	}
	// 创建cookie
	cookie := http.Cookie{
		Name:     dao.CookieKEY,
		Value:    sessionID.String(),
		HttpOnly: true,
		MaxAge:   CookieMaxAge,
	}
	return &cookie, err
}

// CheckMessage 校验邮箱验证码
func CheckMessage(message *model.Message) bool {
	checkMessage, err := dao.CheckMessage(message)
	if err != nil || checkMessage == nil {
		return false
	}
	err = dao.DeleteMessage(message.Email)
	if err != nil {
		return false
	}
	return true
}

// CheckCode 校验邀请码
func CheckCode(codeStr string) (*model.Code, bool) {
	code, err := dao.CheckCode(codeStr)
	if err != nil || code == nil {
		return nil, false
	}
	// 一次性邀请码
	err = dao.DeleteCode(codeStr)
	if err != nil {
		log.Println(err)
	}
	return code, true
}

// InsertUser 保存用户
func InsertUser(user *model.User) {
	err := dao.InsertUser(user)
	if err != nil {
		log.Println(err)
		return
	}
}

// CheckUserName 验证用户名
func CheckUserName(userName string) bool {
	_, err := dao.IsNameOK(userName)
	if err != nil { // 没查到说明可用
		return true
	}
	return false
}

// SendMessage 发送验证码
func SendMessage(message *model.Message) bool {
	err := dao.AddAndUpdateMessage(message)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// IsLegal 判断用户名和密码合法性
func IsLegal(str string) bool {
	if len(str) == 0 || len(str) > 15 {
		return false
	}
	return true
}

// EmailOK 邮箱验证
func EmailOK(email string) bool {
	user, err := dao.GetUserByEmail(email)
	if err != nil || user.ID <= 0 {
		return true
	}
	return false
}

func UpdateUser(user *model.User) error {
	return dao.UpdateUser(user)
}

func GetUserByID(userID int) (*model.User, error) {
	return dao.GetUserByID(userID)
}

func UpdateImg(user *model.User) error {
	return dao.UpdateImg(user)
}

func UpdateEmail(user *model.User) error {
	return dao.UpdateEmail(user)
}
