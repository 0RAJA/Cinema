package server

import (
	"dao"
	"model"
)

// GetPageMovies 获取分页电影信息
func GetPageMovies(pageNo int) (*model.Page, error) {
	page, err := dao.GetPageMovies(1)
	if err != nil {
		return nil, err
	}
	if pageNo > page.TotalPageNo {
		pageNo = page.TotalPageNo
	}
	page, err = dao.GetPageMovies(pageNo)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// GetPageMoviesByName 根据电影名获取分页电影信息
func GetPageMoviesByName(pageNo int, movieName string) (*model.Page, error) {
	page, err := dao.GetPageMoviesByName(1, movieName)
	if err != nil {
		return nil, err
	}
	if pageNo > page.TotalPageNo {
		pageNo = page.TotalPageNo
	}
	page, err = dao.GetPageMoviesByName(1, movieName)
	if err != nil {
		return nil, err
	}
	return page, nil
}

// AddOrUpdateMovie 新增或者更新电影
func AddOrUpdateMovie(movie *model.Movie) error {
	if movie.ID > 0 {
		err := dao.UpdateMovie(movie)
		if err != nil {
			return err
		}
	} else {
		err := dao.AddMovie(movie)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetMovieByID 获取电影信息
func GetMovieByID(movieID int) (*model.Movie, error) {
	return dao.GetMovieByID(movieID)
}

// DeleteMovieByID 删除电影
func DeleteMovieByID(movieID int) error {
	return dao.DeleteMovieByID(movieID)
}

func GetScoreByID(movieID int) (int, error) {
	return dao.GetScoreByID(movieID)
}

func UpdateScore(movieID, score int) error {
	return dao.UpdateScore(movieID, score)
}
