package dao

import (
	"model"
	"utils"
)

// AddTickets 增加票
func AddTickets(tickets []*model.Ticket) error {
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
	sql := "delete from ticket where plan_id = ?;"
	_, err := utils.DB.Exec(sql, planID)
	if err != nil {
		return err
	}
	return err
}

// DeleteTicketsByScreenID 通过ScreenID删除所有相关ticket
func DeleteTicketsByScreenID(ScreenID int) error {
	sql := "delete from ticket where screen_id = ?;"
	_, err := utils.DB.Exec(sql, ScreenID)
	if err != nil {
		return err
	}
	return err
}

// UpdateTicket 通过ticketID更新票信息
func UpdateTicket(ticket *model.Ticket) error {
	sql := "update ticket set state = ?,user_id = ? where id = ?"
	_, err := utils.DB.Exec(sql, ticket.State, ticket.UserID, ticket.ID)
	if err != nil {
		return err
	}
	return err
}

// GetTicketsByPlanID 通过planID获取票
func GetTicketsByPlanID(planID int) ([]*model.Ticket, error) {
	sql := "select id, screen_id, movie_id, plan_id, row, col, state, user_id from ticket;"
	rows, err := utils.DB.Query(sql)
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
			user, err := GetUserByID(ticket.ID)
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