package main

import (
	"controller"
	"net/http"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/views/static"))))
	//去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/views/pages"))))
	//找图片
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("src/views/img"))))
	//去登录
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("/logout", controller.Logout)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//去主页
	http.HandleFunc("/main", controller.Main)
	//验证用户名
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//发送验证码
	http.HandleFunc("/sendMessage", controller.SendMessage)
	//管理员主页
	http.HandleFunc("/manage/main", controller.ManageMain)
	//管理电影
	http.HandleFunc("/manage/movie", controller.ManageMovie)
	//去添加或更新电影
	http.HandleFunc("/manage/toAddOrUpdateMovie", controller.ToAddOrUpdateMovie)
	//新增或更新电影
	http.HandleFunc("/manage/toAddOrUpdateMovie/addOrUpdateMovie", controller.AddOrUpdateMovie)
	//删除电影
	http.HandleFunc("/manage/deleteMovie", controller.DeleteMovie)
	//管理影厅
	http.HandleFunc("/manage/screen", controller.ManageScreen)
	//去新增影厅
	http.HandleFunc("/manage/toAddScreen", controller.ToAddScreen)
	//新增影厅
	http.HandleFunc("/manage/addScreen", controller.AddScreen)
	//删除影厅
	http.HandleFunc("/manage/deleteScreen", controller.DeleteScreen)
	//去修改影厅座位状态
	http.HandleFunc("/manage/toUpdateSeat", controller.ToUpdateSeat)
	//获取座位状态信息
	http.HandleFunc("/manage/getSeatsByScreenID", controller.GetSeatsByScreenID)
	//修改座位状态
	http.HandleFunc("/manage/updateSeats", controller.UpdateSeats)
	//管理剧目
	http.HandleFunc("/manage/plan", controller.ManagePlan)
	//监听
	http.ListenAndServe(":8080", nil)
}
