package server

import (
	"dao"
	"fmt"
	"model"
	"time"
)

func GetTicketsByPlanID(planID int) ([]*model.Ticket, error) {
	return dao.GetTicketsByPlanID(planID)
}

func UpdateTicket(ticket *model.Ticket) error {
	return dao.UpdateTicket(ticket)
}

// IsUnSold 判断是否被卖出去
func IsUnSold(ticketID int) (*model.Ticket, bool) {
	ticket, err := dao.GetTicketByID(ticketID)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return ticket, ticket.State == model.TicketUnsold
}

func GetTicketsByUserID(userID int) ([]*model.Ticket, error) {
	tickets, err := dao.GetTicketsByUserID(userID)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tickets); i++ {
		uTime, _ := time.Parse(model.PlanTimeFormat, tickets[i].UpTime)
		if uTime.Before(time.Now()) {
			tickets[i].IsTimeOut = true
		}
	}
	return tickets, nil
}

func ReturnTicket(ticketID int) error {
	ticket, err := dao.GetTicketByID(ticketID)
	if err != nil {
		return err
	}
	ticket.State = model.TicketUnsold
	err = UpdateTicket(ticket)
	if err != nil {
		return err
	}
	err = UpdateBox(-ticket.Price, ticket.MovieID)
	if err != nil {
		return err
	}
	return nil
}
