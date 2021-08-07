package server

import (
	"dao"
	"model"
)

func GetTicketsByPlanID(planID int) ([]*model.Ticket, error) {
	return dao.GetTicketsByPlanID(planID)
}

func UpdateTicket(ticket *model.Ticket) error {
	return dao.UpdateTicket(ticket)
}
