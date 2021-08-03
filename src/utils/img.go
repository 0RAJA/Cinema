package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	PATH   = "src/views/img/"
	USER   = "user/"
	MOVIE  = "movie/"
	SCREEN = "screen/"
)
const ImgTimeFormat = "2006-1-2-15-04-05-"

// SaveImg 保存图片并返回src
func SaveImg(w http.ResponseWriter, r *http.Request, path string) (src string) {
	f, h, err := r.FormFile("imgPath")
	if err != nil {
		log.Println(f, h, err)
		return ""
	}
	defer f.Close()
	fileName := time.Now().Format(ImgTimeFormat) + h.Filename
	t, err := os.Create(PATH + path + fileName)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	defer t.Close()
	if _, err := io.Copy(t, f); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	return "/img/" + path + fileName
}
