package services

import (
	"goWebService/pkg/models"
	"goWebService/pkg/utils"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(user *models.User) error {
	return utils.DB.Create(user).Error
}

func (u *UserService) GetAllUsers(users *[]models.User) error {
	return utils.DB.Find(users).Error
}

func (u *UserService) GetUser(user *models.User, id string) error {
	return utils.DB.First(user, id).Error
}

func (u *UserService) GetUserByEmail(user *models.User, email string) error {
	return utils.DB.Where("email = ?", email).First(user).Error
}

func (u *UserService) UpdateUser(user *models.User) error {
	return utils.DB.Save(user).Error
}

func (u *UserService) DeleteUser(user *models.User, id string) error {
	return utils.DB.Delete(user, id).Error
}
