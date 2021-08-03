package dao

import (
	"fmt"
	"github.com/gofrs/uuid"
	"model"
	"testing"
)

func TestAddAndUpdateSession(t *testing.T) {
	userID := 1
	user, err := GetUserByID(userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	sessionID, _ := uuid.NewV4()
	session := model.Session{
		ID:       sessionID.String(),
		UserID:   user.ID,
		UserName: user.Name,
		Root:     user.Root,
	}
	err = AddAndUpdateSession(&session)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestDeleteSession(t *testing.T) {
	sessionID := "09820b6d-9a3a-4fee-8ab0-abdec6a48a94"
	err := DeleteSession(sessionID)
	if err != nil {
		fmt.Println(err)
	}
}
