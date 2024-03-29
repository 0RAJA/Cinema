package dao

import (
	"time"

	"cinema/model"
	"cinema/utils"
)

// AddPlan 增加演出计划
func AddPlan(plan *model.Plan) error {
	defer rwmutexPlan.Unlock()
	rwmutexPlan.Lock()
	sql := "insert into plan (screen_id, movie_id, up_time, down_time, price) values (?,?,?,?,?);"
	_, err := utils.DB.Exec(sql, plan.ScreenID, plan.MovieID, plan.UpTime, plan.DownTime, plan.Price)
	if err != nil {
		return err
	}
	sql2 := "SELECT max(id) from plan"
	err = utils.DB.QueryRow(sql2).Scan(&plan.ID)
	if err != nil {
		return err
	}
	// 再增加票
	seats, err := GetSeatsByScreenID(plan.ScreenID)
	if err != nil {
		return err
	}
	screen, err := GetScreenByID(plan.ScreenID)
	if err != nil {
		return err
	}
	movie, err := GetMovieByID(plan.MovieID)
	if err != nil {
		return err
	}
	var tickets []*model.Ticket
	for i := 0; i < len(seats); i++ {
		ticket := model.Ticket{
			ScreenID:   screen.ID,
			MovieID:    movie.ID,
			PlanID:     plan.ID,
			Row:        seats[i].Row,
			Col:        seats[i].Col,
			UserID:     0,
			ScreenName: screen.Name,
			MovieName:  movie.Name,
			UserName:   "",
			UpTime:     plan.UpTime,
			DownTime:   plan.DownTime,
			Price:      plan.Price,
		}
		switch seats[i].State {
		case model.SeatOK:
			ticket.State = model.TicketUnsold
		case model.SeatError:
			ticket.State = model.TicketError
		case model.SeatNone:
			ticket.State = model.TicketNone
		}
		tickets = append(tickets, &ticket)
	}
	err = AddTickets(tickets)
	if err != nil {
		return err
	}
	return err
}

// DeletePlanByID 删除计划
func DeletePlanByID(planID int) error {
	defer rwmutexPlan.Unlock()
	rwmutexPlan.Lock()
	// 先删除票
	err := DeleteTicketsByPlanID(planID)
	if err != nil {
		return err
	}
	sql := "delete from plan where id = ?;"
	_, err = utils.DB.Exec(sql, planID)
	if err != nil {
		return err
	}
	return err
}

// DeletePlansByScreenID 根据screenID删除所有plans
func DeletePlansByScreenID(screenID int) error {
	defer rwmutexPlan.Unlock()
	rwmutexPlan.Lock()
	// 先删除票
	err := DeleteTicketsByScreenID(screenID)
	if err != nil {
		return err
	}
	sql := "delete from plan where screen_id = ?"
	_, err = utils.DB.Exec(sql, screenID)
	if err != nil {
		return err
	}
	return err
}

// GetPagePlans 获取分页计划
func GetPagePlans(pageNo int) (*model.Page, error) {
	defer rwmutexPlan.RUnlock()
	rwmutexPlan.RLock()
	page := model.Page{PageSize: PAGESIZE, PageNo: pageNo}
	// 获取总记录数和总页数
	sql := "select count(*) from plan;"
	err := utils.DB.QueryRow(sql).Scan(&page.TotalRecord)
	if err != nil {
		return nil, err
	}
	page.TotalPageNo = page.TotalRecord / page.PageSize
	if page.TotalRecord%page.PageSize != 0 {
		page.TotalPageNo++
	}
	// 通过limit获取图书
	sql = "select id, screen_id, movie_id, up_time, down_time, price from plan limit ?,?;"
	rows, err := utils.DB.Query(sql, (page.PageNo-1)*page.PageSize, page.PageSize)
	if err != nil {
		return nil, err
	}
	var plans []*model.Plan
	for rows.Next() {
		var plan model.Plan
		err := rows.Scan(&plan.ID, &plan.ScreenID, &plan.MovieID, &plan.UpTime, &plan.DownTime, &plan.Price)
		if err != nil {
			return nil, err
		}
		// 填充信息
		screen, err := GetScreenByID(plan.ScreenID)
		if err != nil {
			return nil, err
		}
		plan.ScreenName = screen.Name
		movie, err := GetMovieByID(plan.MovieID)
		if err != nil {
			return nil, err
		}
		plan.MovieName = movie.Name
		formatTime(&plan)
		plans = append(plans, &plan)
	}
	page.Plans = plans
	return &page, err
}

// GetPlanByID 通过计划ID获取计划
func GetPlanByID(planID int) (*model.Plan, error) {
	defer rwmutexPlan.RUnlock()
	rwmutexPlan.RLock()
	sql := "select id, screen_id, movie_id, up_time, down_time, price from plan where id = ?;"
	var plan model.Plan
	err := utils.DB.QueryRow(sql, planID).Scan(&plan.ID, &plan.ScreenID, &plan.MovieID, &plan.UpTime, &plan.DownTime, &plan.Price)
	if err != nil {
		return nil, err
	}
	screen, err := GetScreenByID(plan.ScreenID)
	if err != nil {
		return nil, err
	}
	plan.ScreenName = screen.Name
	movie, err := GetMovieByID(plan.MovieID)
	if err != nil {
		return nil, err
	}
	plan.MovieName = movie.Name
	formatTime(&plan)
	return &plan, err
}

// UpdatePlan 更新时间还有价格
func UpdatePlan(plan *model.Plan) error {
	defer rwmutexPlan.Unlock()
	rwmutexPlan.Lock()
	sql := "update plan set up_time = ?,down_time = ?,price = ? where id = ?"
	_, err := utils.DB.Exec(sql, plan.UpTime, plan.DownTime, plan.Price, plan.ID)
	if err != nil {
		return err
	}
	return err
}

// GetAllPlans 获取所有plans
func GetAllPlans() ([]*model.Plan, error) {
	defer rwmutexPlan.RUnlock()
	rwmutexPlan.RLock()
	sql := "select id, screen_id, movie_id, up_time, down_time, price from plan"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var plans []*model.Plan
	for rows.Next() {
		var plan model.Plan
		err = rows.Scan(&plan.ID, &plan.ScreenID, &plan.MovieID, &plan.UpTime, &plan.DownTime, &plan.Price)
		if err != nil {
			return nil, err
		}
		formatTime(&plan)
		plans = append(plans, &plan)
	}
	return plans, err
}

// 一键填充并格式化时间
func formatTime(plan *model.Plan) {
	t, _ := time.Parse(model.PlanTimeFormat, plan.UpTime)
	plan.UpTimeFormat = t.Format(model.PlanTimeFormatOK)
	t, _ = time.Parse(model.PlanTimeFormat, plan.DownTime)
	plan.DownTimeFormat = t.Format(model.PlanTimeFormatOK)
}

// GetPlansByScreenAndMovie 通过影厅和电影筛选计划
func GetPlansByScreenAndMovie(screenID, movieID int) ([]*model.Plan, error) {
	defer rwmutexPlan.RUnlock()
	rwmutexPlan.RLock()
	sql := "select id, screen_id, movie_id, up_time, down_time, price from plan where screen_id = ? and movie_id = ?"
	rows, err := utils.DB.Query(sql, screenID, movieID)
	if err != nil {
		return nil, err
	}
	var plans []*model.Plan
	for rows.Next() {
		var plan model.Plan
		err = rows.Scan(&plan.ID, &plan.ScreenID, &plan.MovieID, &plan.UpTime, &plan.DownTime, &plan.Price)
		if err != nil {
			return nil, err
		}
		formatTime(&plan)
		plans = append(plans, &plan)
	}
	return plans, err
}

// GetPlansByMovie 通过电影名筛选演出计划
func GetPlansByMovie(movieID int) ([]*model.Plan, error) {
	defer rwmutexPlan.RUnlock()
	rwmutexPlan.RLock()
	sql := "select p.id, p.screen_id, p.movie_id, p.up_time, p.down_time, p.price,m.name,s.name from plan p join movie m on p.movie_id = m.id join screen s on s.id = p.screen_id where movie_id = ?"
	rows, err := utils.DB.Query(sql, movieID)
	if err != nil {
		return nil, err
	}
	var plans []*model.Plan
	for rows.Next() {
		var plan model.Plan
		err = rows.Scan(&plan.ID, &plan.ScreenID, &plan.MovieID, &plan.UpTime, &plan.DownTime, &plan.Price, &plan.MovieName, &plan.ScreenName)
		if err != nil {
			return nil, err
		}
		formatTime(&plan)
		plans = append(plans, &plan)
	}
	return plans, err
}
