package services

import (
	"coursemicro/models"
	"coursemicro/utils"
)

type usersService struct {
}

var (
	UsersService usersService //what we are saying here is UsersService is of type usersService
)

func (u *usersService) GetUser(userId int64) (*models.User, *utils.ApplicationError) {
	// return models.UserDao.GetUser(userId)
	user, err := models.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
