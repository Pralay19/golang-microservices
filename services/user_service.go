package services

import (
	"coursemicro/models"
	"coursemicro/utils"
)

func GetUser(userId int64) (*models.User, *utils.ApplicationError) {
	return models.GetUser(userId)
}
