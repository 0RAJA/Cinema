package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"model"
	"net/http"
	"server"
	"strconv"
	"text/template"
	"time"
)

// ClientMain 顾客首页
func ClientMain(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/client/client.html"))
	_ = t.Execute(w, "")
}

// GetAllBoxes 获取票房
func GetAllBoxes(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	boxes, err := server.GetAllBoxes()
	if err != nil {
		log.Println(err)
		return
	}
	output, err := json.Marshal(&boxes)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// ToTakeTickets 去选座界面
func ToTakeTickets(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	planID, err := strconv.Atoi(r.FormValue("planID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plan, err := server.GetPlanByID(planID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/client/tickets.html"))
	_ = t.Execute(w, plan)
}

// GetTickets 获取票||选座位
func GetTickets(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	planID, err := strconv.Atoi(r.FormValue("planID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tickets, err := server.GetTicketsByPlanID(planID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := json.Marshal(&tickets)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// GetScreen 获取影厅信息
func GetScreen(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	screenID, err := strconv.Atoi(r.FormValue("screenID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	screen, err := server.GetScreenByID(screenID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := json.Marshal(screen)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// TakeTickets 卖票
func TakeTickets(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	var tickets []*model.Ticket
	err := json.NewDecoder(r.Body).Decode(&tickets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := 0; i < len(tickets); i++ {
		ticket, ok := server.IsUnSold(tickets[i].ID)
		if ok == false {
			continue
		}
		tickets[i].State = model.TicketSold
		tickets[i].UserID = session.UserID
		err := server.UpdateTicket(tickets[i])
		if err != nil {
			fmt.Println(err)
			_, _ = w.Write([]byte("false"))
			return
		}
		err = server.UpdateBox(ticket.Price, ticket.MovieID)
		if err != nil {
			log.Println(err)
			return
		}
	}
	_, _ = w.Write([]byte("true"))
}

// GetMovieMessage 获取单个电影信息
func GetMovieMessage(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/client/movie_message.html"))
	movie, err := server.GetMovieByID(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = t.Execute(w, movie)
}

// GetPageMovies 获取分页电影
func GetPageMovies(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
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
	output, err := json.Marshal(page)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// GetPageMoviesByName 通过电影名获取分页电影
func GetPageMoviesByName(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	pageNo, err := strconv.Atoi(r.FormValue("pageNo"))
	if err != nil || pageNo <= 0 {
		pageNo = 1
	}
	movieName := r.FormValue("movieName")
	var page *model.Page
	if movieName == "" {
		page, err = server.GetPageMovies(pageNo)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		page, err = server.GetPageMoviesByName(pageNo, movieName)
		if err != nil {
			log.Println(err)
			return
		}
	}
	output, err := json.Marshal(page)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// Scoring 评分
func Scoring(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if err != nil {
		_, _ = w.Write([]byte("false"))
		return
	}
	add, err := strconv.Atoi(r.FormValue("score"))
	if err != nil {
		_, _ = w.Write([]byte("false"))
		return
	}
	score, err := server.GetScoreByID(movieID)
	if err != nil {
		_, _ = w.Write([]byte("false"))
		return
	}
	_ = server.UpdateScore(movieID, score+add)
	_, _ = w.Write([]byte("true"))
}

// GetPlansByMovieIDAndTime 获取演出计划
func GetPlansByMovieIDAndTime(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	startTime, err1 := time.Parse(model.MovieTimeFormat, r.FormValue("sTime"))
	endTime, err2 := time.Parse(model.MovieTimeFormat, r.FormValue("eTime"))
	var plans []*model.Plan
	if err1 != nil || err2 != nil {
		plans, err = server.GetPlansByMovie(movieID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		plans, err = server.GetPlansByMovieAndTime(movieID, startTime, endTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	pplans := struct {
		Plans   []*model.Plan
		MovieID int
	}{Plans: plans, MovieID: movieID}
	t := template.Must(template.ParseFiles("src/views/pages/client/plan.html"))
	_ = t.Execute(w, pplans)
}

// GetMovieByID 通过ID获取单个电影信息
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie, err := server.GetMovieByID(movieID)
	output, err := json.Marshal(movie)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// GetUserTickets 获取用户票信息
func GetUserTickets(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	userID := session.UserID
	tickets, err := server.GetTicketsByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/client/order.html"))
	_ = t.Execute(w, tickets)
}

func ReturnTicket(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootClient {
		Main(w, r)
		return
	}
	ticketID, err := strconv.Atoi(r.FormValue("ticketID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.ReturnTicket(ticketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	GetUserTickets(w, r)
}
