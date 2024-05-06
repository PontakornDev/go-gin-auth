package controllers

import (
	"net/http"

	"github.com/PontakornDev/ginAuth/middleware"
	"github.com/PontakornDev/ginAuth/models"
	v1 "github.com/PontakornDev/ginAuth/repositories/v1"
	"github.com/PontakornDev/ginAuth/utils"
	"github.com/gin-gonic/gin"
)

func UsersEndPoint(router *gin.RouterGroup) {
	route := router.Group("/users")
	route.POST("/register", register)
	route.Use(middleware.AuthMiddleware())
	route.GET("/getAll", getAll)
}

func register(ctx *gin.Context) {
	var req = &models.Users{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorMessage(utils.ErrorObject{
			Title:        "Bind Request is Error",
			ErrorMessage: err.Error(),
		}))
		return
	}
	response, err := v1.RegisterUser(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorMessage(utils.ErrorObject{
			Title:        "Register User Error",
			ErrorMessage: err.Error(),
		}))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title: "Register success",
		Item:  response,
	}))

}

func getAll(ctx *gin.Context) {
	response, err := v1.GetAllUser()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorMessage(utils.ErrorObject{
			Title:        "Get All User Error",
			ErrorMessage: err.Error(),
		}))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title: "Register success",
		Item:  response,
	}))
}
