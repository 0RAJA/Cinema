package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"cinema/model"
	"cinema/server"
	"cinema/utils"
)

// ManageMain 管理员主页
func ManageMain(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/manager.html"))
	_ = t.Execute(w, "")
}

// ManageScreen 管理影厅
func ManageScreen(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	pageNo, err := strconv.Atoi(r.FormValue("pageNo"))
	if err != nil || pageNo <= 0 {
		pageNo = 1
	}
	page, err := server.GetPageScreens(pageNo)
	if err != nil {
		log.Println(err)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/screen.html"))
	_ = t.Execute(w, page)
}

// ToAddScreen 去新增影厅
func ToAddScreen(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/screen_add.html"))
	_ = t.Execute(w, r)
}

// AddScreen 增加影厅
func AddScreen(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	name := r.FormValue("name")
	rows, err := strconv.Atoi(r.FormValue("rows"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cols, err := strconv.Atoi(r.FormValue("cols"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if rows <= 0 || cols <= 0 {
		http.Error(w, "数据范围有误", http.StatusBadRequest)
		return
	}
	imgPath := utils.SaveImg(w, r, utils.SCREEN)
	if imgPath == "" {
		imgPath = "/img/screen/1.png"
	}
	screen := model.Screen{
		Name:    name,
		Rows:    rows,
		Cols:    cols,
		ImgPath: imgPath,
	}
	err = server.AddScreen(&screen)
	if err != nil {
		fmt.Println(err)
		return
	}
	ManageScreen(w, r)
}

// DeleteScreen 删除影厅
func DeleteScreen(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	screenID, err := strconv.Atoi(r.FormValue("screenID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = server.DeleteScreenByID(screenID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ManageScreen(w, r)
}

// ToUpdateSeat 去更新座位状态
func ToUpdateSeat(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	screenID, err := strconv.Atoi(r.FormValue("screenID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/seat_update.html"))
	_ = t.Execute(w, screenID)
}

// GetSeatsByScreenID 获取座位
func GetSeatsByScreenID(w http.ResponseWriter, r *http.Request) {
	screenID, err := strconv.Atoi(r.FormValue("screenID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	screenSeats, err := server.GetSeatsByScreenID(screenID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := json.Marshal(&screenSeats)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// UpdateSeats 更新座位状态
func UpdateSeats(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	var seatStates []*model.SeatState
	err := json.NewDecoder(r.Body).Decode(&seatStates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.UpdateSeats(seatStates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}

// ManagePlan 管理演出计划
func ManagePlan(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	pageNo, err := strconv.Atoi(r.FormValue("pageNo"))
	if err != nil || pageNo <= 0 {
		pageNo = 1
	}
	page, err := server.GetPagePlans(pageNo)
	if err != nil {
		log.Println(err)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/plan.html"))
	_ = t.Execute(w, page)
}

// ToAddOrUpdatePlan 去新增或者修改演出计划
func ToAddOrUpdatePlan(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/plan_edit.html"))
	planID, err := strconv.Atoi(r.FormValue("planID"))
	if planID <= 0 {
		_ = t.Execute(w, false)
		return
	}
	plan, err := server.GetPlanByID(planID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = t.Execute(w, plan)
}

// GetMoviesAndScreens 获取电影和影厅信息
func GetMoviesAndScreens(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	movieAndScreen, err := server.GetMoviesAndScreens()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := json.Marshal(&movieAndScreen)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// CheckPlanTime 检查计划时间是否冲突
func CheckPlanTime(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	var planTImeMessage model.PlanTimeMessage
	err := json.NewDecoder(r.Body).Decode(&planTImeMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reply, err := server.CheckPlanTime(&planTImeMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := json.Marshal(reply)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// AddOrUpdatePlan 增加或者修改演出计划
func AddOrUpdatePlan(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	var planMessage model.AddOrUpdatePlanMessage
	err := json.NewDecoder(r.Body).Decode(&planMessage)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.AddOrUpdatePlan(&planMessage)
	w.WriteHeader(200)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("false"))
		return
	}
	_, _ = w.Write([]byte("true"))
}

// DeletePlan 删除计划
func DeletePlan(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	planID, err := strconv.Atoi(r.FormValue("planID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.DeletePlanByID(planID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ManagePlan(w, r)
}

// ManageMovie 管理电影
func ManageMovie(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	pageNo, err := strconv.Atoi(r.FormValue("pageNo"))
	if err != nil || pageNo <= 0 {
		pageNo = 1
	}
	page, err := server.GetPageMovies(pageNo)
	if err != nil {
		log.Println(err)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/movie.html"))
	_ = t.Execute(w, page)
}

// ToAddOrUpdateMovie 去新增或者修改电影
func ToAddOrUpdateMovie(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/manager/movie_edit.html"))
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if movieID <= 0 {
		_ = t.Execute(w, false)
		return
	}
	movie, err := server.GetMovieByID(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = t.Execute(w, movie)
}

// AddOrUpdateMovie 新增电影
func AddOrUpdateMovie(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	movieID, _ := strconv.Atoi(r.FormValue("movieID"))
	length, err := strconv.Atoi(r.FormValue("length"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	imgPath := utils.SaveImg(w, r, utils.MOVIE)
	if imgPath == "" {
		imgPath = "/img/movie/1.png"
	}
	var movie = model.Movie{
		ID:       movieID,
		Name:     r.FormValue("name"),
		UpDate:   r.FormValue("upDate"),
		Director: r.FormValue("director"),
		Type:     r.FormValue("type"),
		Area:     r.FormValue("area"),
		Star:     r.FormValue("star"),
		Length:   length,
		Intro:    r.FormValue("intro"),
		ImgPath:  imgPath,
	}
	// 将请求体中的 JSON 数据解析到结构体中
	// 发生错误，返回400 错误码
	// err := json.NewDecoder(r.Body).Decode(&movie)
	// fmt.Println(movie, err)
	// if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	// }
	err = server.AddOrUpdateMovie(&movie)
	if err != nil {
		log.Println(err)
		return
	}
	ManageMovie(w, r)
}

// DeleteMovie 删除电影
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootManager {
		Main(w, r)
		return
	}
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.DeleteMovieByID(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ManageMovie(w, r)
}
