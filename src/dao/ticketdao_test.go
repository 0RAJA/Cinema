package dao

import (
	"fmt"
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
