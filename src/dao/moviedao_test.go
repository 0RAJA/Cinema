package dao

import (
	"fmt"
	"testing"
	"time"

	"cinema/model"
)

func TestAddMovie(t *testing.T) {
	movie := model.Movie{
		Name:     "泰坦号",
		UpDate:   time.Date(2020, 3, 26, 5, 0, 0, 0, time.Local).Format(model.MovieTimeFormat),
		Director: "1",
		Type:     "1",
		Area:     "1",
		Star:     "1",
		Length:   80,
		Intro:    "222",
		ImgPath:  "/img/movie/1.png",
	}
	err := AddMovie(&movie)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestUpdateMovie(t *testing.T) {
	movie := model.Movie{
		ID:       1,
		Name:     "泰坦尼克号",
		UpDate:   time.Date(2020, 3, 26, 5, 0, 0, 0, time.Local).Format(model.MovieTimeFormat),
		Director: "2",
		Type:     "2",
		Area:     "2",
		Star:     "2",
		Length:   100,
		Intro:    "111",
		ImgPath:  "/img/movie/1.png",
	}
	err := UpdateMovie(&movie)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetAllMovies(T *testing.T) {
	movies, err := GetAllMovies()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range movies {
		fmt.Println(*v)
	}
}

func TestGetMovieByID(t *testing.T) {
	movieID := 1
	movie, err := GetMovieByID(movieID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*movie)
}

func TestGetPageMovies(t *testing.T) {
	pageNo := 1
	page, err := GetPageMovies(pageNo)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range page.Movies {
		fmt.Println(*v)
	}
}
