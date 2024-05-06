package controllers

import (
	"net/http"

	"github.com/PontakornDev/ginAuth/middleware"
	"github.com/PontakornDev/ginAuth/models"
	v1 "github.com/PontakornDev/ginAuth/repositories/v1"
	"github.com/PontakornDev/ginAuth/utils"
	"github.com/gin-gonic/gin"
)

func AuthEndPoint(router *gin.RouterGroup) {
	router.POST("/auth", Login)
}

func Login(ctx *gin.Context) {
	var req = &models.Auth{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	users, err := v1.QueryPasswordByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorMessage(utils.ErrorObject{
			Title:        "Query Password By Username Error",
			ErrorMessage: err.Error(),
		}))
		return
	}
	hashStatus := utils.CheckPasswordHash(req.Password, users.Password)
	if !hashStatus {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorMessage(utils.ErrorObject{
			Title:        "Login is fail",
			ErrorMessage: "password is incorrect",
		}))
		return
	}

	token, err := middleware.GenerateToken(users)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorMessage(utils.ErrorObject{
			Title:        "generate token is fail",
			ErrorMessage: err.Error(),
		}))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title: "Login is success",
		Item: &models.Token{
			Token: token,
		},
	}))
}
