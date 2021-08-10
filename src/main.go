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
	//去bootstrap
	http.Handle("/bootstrap/", http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("src/views/bootstrap"))))
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
	//管理计划
	http.HandleFunc("/manage/plan", controller.ManagePlan)
	//去新增或更新计划
	http.HandleFunc("/manage/toAddOrUpdatePlan", controller.ToAddOrUpdatePlan)
	//获取电影和影厅信息
	http.HandleFunc("/manage/getMoviesAndScreens", controller.GetMoviesAndScreens)
	//检查计划时间是否冲突
	http.HandleFunc("/manage/checkPlanTime", controller.CheckPlanTime)
	//增加或者更新计划
	http.HandleFunc("/manage/addOrUpdatePlan", controller.AddOrUpdatePlan)
	//删除计划
	http.HandleFunc("/manage/deletePlan", controller.DeletePlan)
	//顾客首页
	http.HandleFunc("/client/main", controller.ClientMain)
	//获取票房
	http.HandleFunc("/client/getAllBoxes", controller.GetAllBoxes)
	//获取单个电影信息
	http.HandleFunc("/client/getMovieMessage", controller.GetMovieMessage)
	//获取分页电影信息
	http.HandleFunc("/client/getPageMovies", controller.GetPageMovies)
	//根据电影名获取分页电影信息
	http.HandleFunc("/client/getPageMoviesByName", controller.GetPageMoviesByName)
	//评分
	http.HandleFunc("/client/scoring", controller.Scoring)
	//获取演出计划
	http.HandleFunc("/client/getPlansByMovieIDAndTime", controller.GetPlansByMovieIDAndTime)
	//获取单个电影信息
	http.HandleFunc("/client/getMovieByID", controller.GetMovieByID)
	//去选座界面
	http.HandleFunc("/client/toTakeTickets", controller.ToTakeTickets)
	//获取票
	http.HandleFunc("/client/getTickets", controller.GetTickets)
	//获取影厅信息
	http.HandleFunc("/client/getScreen", controller.GetScreen)
	//卖票
	http.HandleFunc("/client/takeTickets", controller.TakeTickets)
	//订单
	http.HandleFunc("/client/order", controller.GetUserTickets)
	//退票
	http.HandleFunc("/client/returnTicket", controller.ReturnTicket)
	//基本信息
	http.HandleFunc("/profile", controller.Profile)
	//更新用户信息
	http.HandleFunc("/user/updateMessage", controller.UpdateMessage)
	//更新头像
	http.HandleFunc("/user/updateImg", controller.UpdateImg)
	//更新邮箱
	http.HandleFunc("/user/updateEmail", controller.UpdateEmail)
	//获取用户信息
	http.HandleFunc("/user/getUserMessage", controller.GetUserMessage)
	//监听
	_ = http.ListenAndServe(":8080", nil)
}
