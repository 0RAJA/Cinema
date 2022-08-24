package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"cinema/model"
	"cinema/server"
)

// GetAllCodes 获取所有邀请码
func GetAllCodes(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootAdmin {
		Main(w, r)
		return
	}
	codes, err := server.GetAllCodes()
	if err != nil {
		log.Println(err)
		return
	}
	output, err := json.Marshal(codes)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
}

// DeleteCode 删除邀请码
func DeleteCode(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootAdmin {
		Main(w, r)
		return
	}
	code := r.FormValue("code")
	err := server.DeleteCodeByID(code)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t := template.Must(template.ParseFiles("src/views/pages/admin/account.html"))
	_ = t.Execute(w, "")
}

// AddCode 增加新邀请码
func AddCode(w http.ResponseWriter, r *http.Request) {
	session, ok := server.IsLogin(r)
	if ok == false || session == nil || session.Root != model.RootAdmin {
		Main(w, r)
		return
	}
	root, err := strconv.Atoi(r.FormValue("root"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	code, err := server.AddCode(root)
	if err != nil {
		log.Println(err)
		_, _ = w.Write([]byte(""))
		return
	}
	_, _ = w.Write([]byte(code))
}
