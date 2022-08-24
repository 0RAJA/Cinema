package dao

import (
	"cinema/model"
	"cinema/utils"
)

// PAGESIZE 分页大小
const PAGESIZE = 8

// AddMovie 新增电影信息
func AddMovie(movie *model.Movie) error {
	defer rwmutexMovie.Unlock()
	rwmutexMovie.Lock()
	sql := "insert into movie (name, up_date, type, area, director,star, length, intro, img_path) values (?,?,?,?,?,?,?,?,?);"
	_, err := utils.DB.Exec(sql, movie.Name, movie.UpDate, movie.Type, movie.Area, movie.Director, movie.Star, movie.Length, movie.Intro, movie.ImgPath)
	if err != nil {
		return err
	}
	// 增加票房信息
	sql = "select max(id) from movie;"
	var movieID int
	err = utils.DB.QueryRow(sql).Scan(&movieID)
	if err != nil {
		return err
	}
	err = AddBox(movieID)
	if err != nil {
		return err
	}
	return err
}

// UpdateMovie 更新电影信息
func UpdateMovie(movie *model.Movie) error {
	defer rwmutexMovie.Unlock()
	rwmutexMovie.Lock()
	sql := "update movie set name = ?,up_date = ?,type = ?,area = ?,director = ?,star = ?,length = ?,intro = ?,img_path = ? where id = ?"
	_, err := utils.DB.Exec(sql, movie.Name, movie.UpDate, movie.Type, movie.Area, movie.Director, movie.Star, movie.Length, movie.Intro, movie.ImgPath, movie.ID)
	if err != nil {
		return err
	}
	return err
}

// GetAllMovies 获取所有电影信息
func GetAllMovies() ([]*model.Movie, error) {
	defer rwmutexMovie.RUnlock()
	rwmutexMovie.RLock()
	sql := "select id, name, up_date, type, area, director,star, length, intro, img_path  from movie order by score desc;"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var movies []*model.Movie
	for rows.Next() {
		var movie model.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.UpDate, &movie.Type, &movie.Area, &movie.Director, &movie.Star, &movie.Length, &movie.Intro, &movie.ImgPath)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	return movies, err
}

// GetMovieByID 通过电影ID查询电影信息
func GetMovieByID(movieID int) (*model.Movie, error) {
	defer rwmutexMovie.RUnlock()
	rwmutexMovie.RLock()
	sql := "select id, name, up_date, type, area, director,star, length, intro, img_path,score  from movie where id = ?;"
	var movie model.Movie
	err := utils.DB.QueryRow(sql, movieID).Scan(&movie.ID, &movie.Name, &movie.UpDate, &movie.Type, &movie.Area, &movie.Director, &movie.Star, &movie.Length, &movie.Intro, &movie.ImgPath, &movie.Score)
	if err != nil {
		return nil, err
	}
	return &movie, err
}

// GetPageMovies 通过pageNo获取当前页数电影信息
func GetPageMovies(pageNo int) (*model.Page, error) {
	defer rwmutexMovie.RUnlock()
	rwmutexMovie.RLock()
	page := model.Page{PageSize: PAGESIZE, PageNo: pageNo}
	// 获取总记录数和总页数
	sql := "select count(*) from movie;"
	err := utils.DB.QueryRow(sql).Scan(&page.TotalRecord)
	if err != nil {
		return nil, err
	}
	page.TotalPageNo = page.TotalRecord / page.PageSize
	if page.TotalRecord%page.PageSize != 0 {
		page.TotalPageNo++
	}
	// 通过limit获取图书
	sql = "select * from movie limit ?,?;"
	rows, err := utils.DB.Query(sql, (page.PageNo-1)*page.PageSize, page.PageSize)
	if err != nil {
		return nil, err
	}
	var movies []*model.Movie
	for rows.Next() {
		var movie model.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.UpDate, &movie.Type, &movie.Area, &movie.Director, &movie.Star, &movie.Length, &movie.Intro, &movie.ImgPath, &movie.Score)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	page.Movies = movies
	return &page, err
}

// GetPageMoviesByName 通过pageNo获取当前页数电影信息
func GetPageMoviesByName(pageNo int, movieName string) (*model.Page, error) {
	defer rwmutexMovie.RUnlock()
	rwmutexMovie.RLock()
	page := model.Page{PageSize: PAGESIZE, PageNo: pageNo}
	limit := "%" + movieName + "%"
	// 获取总记录数和总页数
	sql := "select count(*) from movie where name like ?;"
	err := utils.DB.QueryRow(sql, limit).Scan(&page.TotalRecord)
	if err != nil {
		return nil, err
	}
	page.TotalPageNo = page.TotalRecord / page.PageSize
	if page.TotalRecord%page.PageSize != 0 {
		page.TotalPageNo++
	}
	// 通过limit获取图书
	sql = "select * from movie where name like ? limit ?,?;"
	rows, err := utils.DB.Query(sql, limit, (page.PageNo-1)*page.PageSize, page.PageSize)
	if err != nil {
		return nil, err
	}
	var movies []*model.Movie
	for rows.Next() {
		var movie model.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.UpDate, &movie.Type, &movie.Area, &movie.Director, &movie.Star, &movie.Length, &movie.Intro, &movie.ImgPath, &movie.Score)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}
	page.Movies = movies
	return &page, err
}

// DeleteMovieByID 删除电影信息
func DeleteMovieByID(movieID int) error {
	defer rwmutexMovie.Unlock()
	rwmutexMovie.Lock()
	sql := "delete from movie where id = ?;"
	_, err := utils.DB.Exec(sql, movieID)
	if err != nil {
		return err
	}
	return err
}

// GetScoreByID 获取分数
func GetScoreByID(movieID int) (int, error) {
	defer rwmutexMovie.RUnlock()
	rwmutexMovie.RLock()
	sql := "select score from movie where id = ?;"
	var score int
	err := utils.DB.QueryRow(sql, movieID).Scan(&score)
	if err != nil {
		return 0, err
	}
	return score, err
}

// UpdateScore 更新分数
func UpdateScore(movieID, score int) error {
	defer rwmutexMovie.Unlock()
	rwmutexMovie.Lock()
	sql := "update movie set score = ? where id = ?"
	_, err := utils.DB.Exec(sql, score, movieID)
	if err != nil {
		return err
	}
	return nil
}
