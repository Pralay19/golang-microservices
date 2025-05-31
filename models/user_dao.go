package models

import (
	"coursemicro/utils"
	"fmt"
	"net/http"
)

// DATABASE
var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Pralay",
			LastName:  "Saw",
			Email:     "myemail@hotmail.com",
		},
		100: {
			Id:        100,
			FirstName: "Rajan",
			LastName:  "Khade",
			Email:     "youremail@hotmail.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		// return nil, errors.New(fmt.Sprintf("user %v was not found", userId)) - this was shown to be used but the below was recommended by Copilot.
		// return nil, fmt.Errorf("user %v was not found", userId)
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
