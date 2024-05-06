package v1

import (
	"github.com/PontakornDev/ginAuth/models"
	"github.com/PontakornDev/ginAuth/repositories"
	"github.com/PontakornDev/ginAuth/utils"
)

func RegisterUser(r *models.Users) (*models.Users, error) {
	hash, err := utils.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}
	r.Password = hash
	if err := repositories.DB.Create(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func QueryPasswordByUsername(r *models.Auth) (*models.Users, error) {
	users := models.Users{}
	if err := repositories.DB.Find(&users).Where("username = ?", r.Username).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func GetAllUser() (*[]models.Users, error) {
	users := []models.Users{}
	if err := repositories.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
