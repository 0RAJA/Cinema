package dao

import (
	"fmt"
	"log"
	"testing"
)

func TestGetTicketsByPlanID(t *testing.T) {
	planID := 9
	tickets, err := GetTicketsByPlanID(planID)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(tickets); i++ {
		fmt.Println(tickets[i])
	}
}

func TestGetTicketsByUserID(T *testing.T) {
	userID := 13
	tickets, err := GetTicketsByUserID(userID)
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range tickets {
		fmt.Println(v)
	}
}
