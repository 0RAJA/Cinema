package server

import (
	"dao"
	"log"
	"model"
	"time"
)

// GetPagePlans 获取分页计划信息
func GetPagePlans(pageNo int) (*model.Page, error) {
	page, err := dao.GetPagePlans(1)
	if err != nil {
		return nil, err
	}
	if pageNo > page.TotalPageNo {
		pageNo = page.TotalPageNo
	}
	page, err = dao.GetPagePlans(pageNo)
	if err != nil {
		return nil, err
	}
	return page, err
}

// AddOrUpdatePlan 增加或者更新Plan
func AddOrUpdatePlan(planMessage *model.AddOrUpdatePlanMessage) error {
	upTimeFormat, _ := time.Parse(model.PlanTimeFormatOK, planMessage.UpTimeFormat)
	upTime := upTimeFormat.Format(model.PlanTimeFormat)
	downTimeFormat, _ := time.Parse(model.PlanTimeFormatOK, planMessage.DownTimeFormat)
	downTime := downTimeFormat.Format(model.PlanTimeFormat)
	var plan = model.Plan{
		ScreenID: planMessage.ScreenID,
		MovieID:  planMessage.MovieID,
		UpTime:   upTime,
		DownTime: downTime,
		Price:    planMessage.Price,
	}
	if planMessage.ID > 0 {
		plan.ID = planMessage.ID
		err := dao.UpdatePlan(&plan)
		if err != nil {
			return err
		}
	} else {
		err := dao.AddPlan(&plan)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPlanByID 根据ID获取Plan信息
func GetPlanByID(planID int) (*model.Plan, error) {
	return dao.GetPlanByID(planID)
}

// DeletePlanByID 根据ID删除plan信息
func DeletePlanByID(planID int) error {
	return dao.DeletePlanByID(planID)
}

// GetMoviesAndScreens 获取电影和影厅信息
func GetMoviesAndScreens() (*model.MovieAndScreen, error) {
	movies, err := dao.GetAllMovies()
	if err != nil {
		return nil, err
	}
	screens, err := dao.GetAllScreens()
	if err != nil {
		return nil, err
	}
	movieAndScreen := model.MovieAndScreen{
		Screens: screens,
		Movies:  movies,
	}
	return &movieAndScreen, nil
}

// CheckPlanTime 检查时间合法性
func CheckPlanTime(message *model.PlanTimeMessage) (*model.ReplyTimeMessage, error) {
	plans, err := dao.GetPlansByScreenAndMovie(message.ScreenID, message.MovieID)
	if err != nil {
		log.Println(err)
		return &model.ReplyTimeMessage{}, nil
	}
	startTime, _ := time.Parse(model.PlanTimeFormatOK, message.UpTimeFormat)
	movie, err := GetMovieByID(message.MovieID)
	if err != nil {
		log.Println(err)
		return &model.ReplyTimeMessage{}, nil
	}
	endTime := startTime.Add(time.Duration(movie.Length) * time.Minute)
	reply := model.ReplyTimeMessage{
		DownTimeFormat: endTime.Format(model.PlanTimeFormatOK),
		IsOK:           false,
	}
	for i := 0; i < len(plans); i++ {
		sTime, _ := time.Parse(model.PlanTimeFormat, plans[i].UpTime)
		eTime, _ := time.Parse(model.PlanTimeFormat, plans[i].DownTime)
		if sTime.After(endTime) || eTime.Before(startTime) {
			continue
		}
		reply.IsOK = false
		return &reply, nil
	}
	reply.IsOK = true
	return &reply, nil
}
