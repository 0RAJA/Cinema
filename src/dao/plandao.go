package dao

import (
	"model"
	"utils"
)

// AddPlan 增加演出计划
func AddPlan(plan *model.Plan) error {
	sql := "insert into plan (screen_id, movie_id, up_time, down_time, price) values (?,?,?,?,?);"
	_, err := utils.DB.Exec(sql, plan.ScreenID, plan.MovieID, plan.UpTime, plan.DownTime, plan.Price)
	if err != nil {
		return err
	}
	//再增加票
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
			State:      seats[i].State,
			UserID:     0,
			ScreenName: screen.Name,
			MovieName:  movie.Name,
			UserName:   "",
			UpTime:     plan.UpTime,
			DownTime:   plan.DownTime,
			Price:      plan.Price,
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
	//先删除票
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
	//先删除票
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

// GetPagesPlans 获取分页计划
func GetPagesPlans(pageNo int, planCondition *model.PlanCondition) (*model.Page, error) {
	page := model.Page{PageSize: PAGESIZE, PageNo: pageNo}
	//获取总记录数和总页数
	sql := "select count(*) from plan where screen_id = ? and movie_id = ? and up_time > ?;"
	err := utils.DB.QueryRow(sql, planCondition.ScreenID, planCondition.MovieID, planCondition.UpTime).Scan(&page.TotalRecord)
	if err != nil {
		return nil, err
	}
	page.TotalPageNo = page.TotalRecord / page.PageSize
	if page.TotalRecord%page.PageSize != 0 {
		page.TotalPageNo++
	}
	//通过limit获取图书
	sql = "select id, screen_id, movie_id, up_time, down_time, price from plan where screen_id = ? and movie_id = ? and up_time > ? limit ?,?;"
	rows, err := utils.DB.Query(sql, planCondition.ScreenID, planCondition.MovieID, planCondition.UpTime, (page.PageNo-1)*page.PageSize, page.PageSize)
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
		//填充信息
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
		plans = append(plans, &plan)
	}
	page.Plans = plans
	return &page, err
}

// GetPlanByID 通过计划ID获取计划
func GetPlanByID(planID int) (*model.Plan, error) {
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
	return &plan, err
}
