package dao

import (
	"cinema/model"
	"cinema/utils"
)

// AddTickets 增加票
func AddTickets(tickets []*model.Ticket) error {
	defer rwmutexTicket.Unlock()
	rwmutexTicket.Lock()
	sql := "insert into ticket (screen_id, movie_id, plan_id, row, col, state, user_id) values (?,?,?,?,?,?,?);"
	for i := range tickets {
		_, err := utils.DB.Exec(sql, tickets[i].ScreenID, tickets[i].MovieID, tickets[i].PlanID, tickets[i].Row, tickets[i].Col, tickets[i].State, tickets[i].UserID)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteTicketsByPlanID 通过planID删除所有相关ticket
func DeleteTicketsByPlanID(planID int) error {
	defer rwmutexTicket.Unlock()
	rwmutexTicket.Lock()
	sql := "delete from ticket where plan_id = ?;"
	_, err := utils.DB.Exec(sql, planID)
	if err != nil {
		return err
	}
	return err
}

// DeleteTicketsByScreenID 通过ScreenID删除所有相关ticket
func DeleteTicketsByScreenID(ScreenID int) error {
	defer rwmutexTicket.Unlock()
	rwmutexTicket.Lock()
	sql := "delete from ticket where screen_id = ?;"
	_, err := utils.DB.Exec(sql, ScreenID)
	if err != nil {
		return err
	}
	return err
}

// UpdateTicket 通过ticketID更新票信息
func UpdateTicket(ticket *model.Ticket) error {
	defer rwmutexTicket.RUnlock()
	rwmutexTicket.RLock()
	sql := "update ticket set state = ?,user_id = ? where id = ?"
	_, err := utils.DB.Exec(sql, ticket.State, ticket.UserID, ticket.ID)
	if err != nil {
		return err
	}
	return err
}

// GetTicketsByPlanID 通过planID获取票
func GetTicketsByPlanID(planID int) ([]*model.Ticket, error) {
	defer rwmutexTicket.RUnlock()
	rwmutexTicket.RLock()
	sql := "select id, screen_id, movie_id, plan_id, row, col, state, user_id from ticket where plan_id = ?;"
	rows, err := utils.DB.Query(sql, planID)
	if err != nil {
		return nil, err
	}
	var tickets []*model.Ticket
	var plan = &model.Plan{}
	isTrue := true
	for rows.Next() {
		var ticket model.Ticket
		err = rows.Scan(&ticket.ID, &ticket.ScreenID, &ticket.MovieID, &ticket.PlanID, &ticket.Row, &ticket.Col, &ticket.State, &ticket.UserID)
		if err != nil {
			return nil, err
		}
		if isTrue {
			plan, err = GetPlanByID(planID)
			isTrue = false
		}
		if ticket.State == model.TicketSold {
			user, err := GetUserByID(ticket.UserID)
			if err != nil {
				return nil, err
			}
			ticket.UserName = user.Name
		}
		ticket.ScreenName, ticket.MovieName, ticket.UpTime, ticket.DownTime = plan.ScreenName, plan.MovieName, plan.UpTime, plan.DownTime
		tickets = append(tickets, &ticket)
	}
	return tickets, err
}

// GetTicketByID 通过ID查询ticket信息
func GetTicketByID(ticketID int) (*model.Ticket, error) {
	defer rwmutexTicket.RUnlock()
	rwmutexTicket.RLock()
	sql := "select t.id, t.screen_id, t.movie_id, t.plan_id, t.row, t.col, t.state, t.user_id,m.name,s.name,p.up_time,p.down_time,p.price from ticket t join movie m on m.id = t.movie_id join screen s on s.id = t.screen_id join plan p on p.id = t.plan_id where t.id = ?;"
	var ticket model.Ticket
	err := utils.DB.QueryRow(sql, ticketID).Scan(&ticket.ID, &ticket.ScreenID, &ticket.MovieID, &ticket.PlanID, &ticket.Row, &ticket.Col, &ticket.State, &ticket.UserID, &ticket.MovieName, &ticket.ScreenName, &ticket.UpTime, &ticket.DownTime, &ticket.Price)
	if err != nil {
		return nil, err
	}
	return &ticket, err
}

// GetTicketsByUserID 通过UserID获取票
func GetTicketsByUserID(userID int) ([]*model.Ticket, error) {
	defer rwmutexTicket.RUnlock()
	rwmutexTicket.RLock()
	sql := "select t.id, t.screen_id, t.movie_id, t.plan_id, t.row, t.col, t.state, t.user_id,m.name,s.name,p.up_time,p.down_time,p.price,u.name from ticket t join movie m on m.id = t.movie_id join screen s on s.id = t.screen_id join plan p on p.id = t.plan_id join users u on u.id = t.user_id where t.user_id = ? and t.state = ?;"
	var tickets []*model.Ticket
	rows, err := utils.DB.Query(sql, userID, model.TicketSold)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ticket model.Ticket
		err := rows.Scan(&ticket.ID, &ticket.ScreenID, &ticket.MovieID, &ticket.PlanID, &ticket.Row, &ticket.Col, &ticket.State, &ticket.UserID, &ticket.MovieName, &ticket.ScreenName, &ticket.UpTime, &ticket.DownTime, &ticket.Price, &ticket.UserName)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, &ticket)
	}
	return tickets, nil
}
