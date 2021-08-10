package controller

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"server"
	"text/template"
	"utils"
)

// Profile 用户界面
func Profile(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/user/self.html"))
	_ = t.Execute(w, user)
}

// UpdateMessage 更新用户信息
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}

// UpdateImg 更新头像
func UpdateImg(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	imgPath := utils.SaveImg(w, r, utils.USER)
	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ImgPath = imgPath
	err = server.UpdateImg(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Profile(w, r)
}

// UpdateEmail 更新邮箱
func UpdateEmail(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	var message model.Message
	message.Email = r.FormValue("email")
	message.Str = r.FormValue("str")
	ok = server.CheckMessage(&message) //验证验证码
	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Email = message.Email
	err = server.UpdateEmail(user)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("false"))
		return
	}
	_, _ = w.Write([]byte("true"))
}
