package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"cinema/model"
	"cinema/server"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sessions, err := server.GetAllComments(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	output, err := json.Marshal(sessions)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	text := r.FormValue("comment")
	movieID, err := strconv.Atoi(r.FormValue("movieID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := server.GetUserByID(session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment := model.Comment{
		MovieID:  movieID,
		UserID:   session.UserID,
		Cnt:      0,
		Time:     time.Now().Format(model.PlanTimeFormat),
		Text:     text,
		UserName: session.UserName,
		ImgPath:  user.ImgPath,
	}
	err = server.AddComment(&comment)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(200)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	commentID, err := strconv.Atoi(r.FormValue("commentID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err = server.IsSelfComment(commentID, session.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ok == false {
		_, _ = w.Write([]byte("false"))
		return
	}
	err = server.DeleteComment(commentID)
	if err != nil {
		_, _ = w.Write([]byte("false"))
	} else {
		_, _ = w.Write([]byte("true"))
	}
}

func LikeComment(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil {
		Main(w, r)
		return
	}
	commentID, err := strconv.Atoi(r.FormValue("commentID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	add, err := strconv.Atoi(r.FormValue("add"))
	if err != nil || add < -1 || add > 1 {
		_, _ = w.Write([]byte("false"))
		return
	}
	comment, err := server.GetCommentByID(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment.Cnt += add
	err = server.UpdateComment(comment)
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte("false"))
		return
	}
	_, _ = w.Write([]byte("true"))
}
