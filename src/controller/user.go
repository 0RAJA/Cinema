package controller

import (
	"dao"
	"encoding/json"
	"log"
	"model"
	"net/http"
	"server"
	"text/template"
	"utils"
)

// Main 主页
func Main(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	t := template.Must(template.ParseFiles("src/views/index.html"))
	if ok == false || session == nil {
		_ = t.Execute(w, false)
		return
	} else {
		_ = t.Execute(w, true)
		return
	}
}

// Login 登录
func Login(w http.ResponseWriter, r *http.Request) {
	session, flag := server.IsLogin(r)
	if flag == true { //登录过了
		//TODO 去Port
		selectPort(w, session.Root)
		return
	}
	userName, password := r.PostFormValue("username"), r.FormValue("password")
	//没有登录
	user, err := server.CheckUserNamePwd(userName, password)
	if err != nil {
		reply(w, "src/views/pages/user/login.html", "用户名或密码错误")
		return
	} else { //登陆成功
		cookie, err := server.CreatSession(user)
		if err != nil {
			log.Println(err)
			return
		}
		//将cookie发给浏览器
		http.SetCookie(w, cookie)
		//去对应端口
		selectPort(w, user.Root)
	}
}

// reply 响应
func reply(w http.ResponseWriter, filePath, str string) {
	t := template.Must(template.ParseFiles(filePath))
	_ = t.Execute(w, str)
}

// selectPort 给用户分配端口
func selectPort(w http.ResponseWriter, root int) {
	var t = new(template.Template)
	switch root {
	case model.RootClient:
		//TODO 顾客页面
		t = template.Must(template.ParseFiles("src/views/pages/client/client.html"))
	case model.RootManager:
		//TODO 经理页面
		t = template.Must(template.ParseFiles("src/views/pages/manager/manager.html"))
	case model.RootAdmin:
		//TODO 管理员界面
		t = template.Must(template.ParseFiles("src/views/pages/admin/account.html"))
	}
	_ = t.Execute(w, "")
}

// Logout 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	//注销
	//获取Cookie
	cookie, _ := r.Cookie(dao.CookieKEY)
	if cookie != nil {
		cookieValue := cookie.Value
		err := server.DeleteCookie(cookieValue)
		if err != nil {
			log.Panicln(err)
			return
		}
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	//去首页
	Main(w, r)
}

// Regist 注册
func Regist(w http.ResponseWriter, r *http.Request) {
	var message model.Message
	message.Email, message.Str = r.FormValue("email"), r.FormValue("str")
	ok := server.EmailOK(message.Email)
	if ok == false {
		reply(w, "src/views/pages/user/regist.html", "邮箱已被注册")
		return
	}
	ok = server.CheckMessage(&message) //验证验证码
	if ok == false {
		reply(w, "src/views/pages/user/regist.html", "验证码错误")
		return
	}
	userName, password, codeStr := r.FormValue("username"), r.FormValue("password"), r.FormValue("code")
	if server.IsLegal(userName) == false || server.IsLegal(password) == false {
		reply(w, "src/views/pages/user/regist.html", "用户名或密码格式有误")
		return
	}
	var root = 0
	if codeStr != "" { //输入了邀请码
		code, ok := server.CheckCode(codeStr)
		if ok == false {
			reply(w, "src/views/pages/user/regist.html", "邀请码不存在")
			return
		}
		root = code.Root
	}
	imgPath := utils.SaveImg(w, r, utils.USER)
	if imgPath == "" {
		imgPath = "/img/user/1.png" //默认头像
	}
	user := model.User{
		Name:     userName,
		Password: password,
		Email:    message.Email,
		Root:     root,
		ImgPath:  imgPath,
	}
	//注册成功
	server.InsertUser(&user)
	//去主页
	Main(w, r)
}

// CheckUserName 检查用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	if server.CheckUserName(userName) {
		_, _ = w.Write([]byte("true"))
	} else {
		_, _ = w.Write([]byte("false"))
	}
}

// SendMessage 发送验证码
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var message model.Message
	message.Email = r.FormValue("email")
	message.Str = utils.SendEmail(message.Email)
	ok := server.SendMessage(&message)
	if ok == true {
		_, _ = w.Write([]byte("true"))
	} else {
		_, _ = w.Write([]byte("false"))
	}
}

// GetUserMessage 获取用户信息
func GetUserMessage(w http.ResponseWriter, r *http.Request) {
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
	output, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}
