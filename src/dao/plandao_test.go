package dao

import (
	"fmt"
	"model"
	"testing"
	"time"
)

func TestAddPlan(t *testing.T) {
	plan := model.Plan{
		ScreenID: 10,
		MovieID:  2,
		UpTime:   time.Date(2021, 11, 1, 15, 20, 0, 0, time.Local).Format(model.PlanTimeFormat),
		DownTime: time.Date(2021, 11, 1, 15, 20, 0, 0, time.Local).Add(80 * time.Minute).Format(model.PlanTimeFormat),
		Price:    80,
	}
	err := AddPlan(&plan)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDeletePlanByID(t *testing.T) {
	planID := 8
	err := DeletePlanByID(planID)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestGetPagePlans(t *testing.T) {
	page, err := GetPagePlans(1)
	if err != nil {
		return
	}
	fmt.Println(page)
	for i := 0; i < len(page.Plans); i++ {
		fmt.Println(*page.Plans[i])
	}
}
