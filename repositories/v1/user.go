package v1

import (
	"github.com/PontakornDev/ginAuth/models"
	"github.com/PontakornDev/ginAuth/repositories"
	"github.com/PontakornDev/ginAuth/utils"
)

func RegisterUser(r models.Users) error {
	hash, err := utils.HashPassword(r.Password)
	if err != nil {
		return err
	}
	r.Password = hash
	if err := repositories.DB.Create(&r).Error; err != nil {
		return err
	}
	return nil
}
