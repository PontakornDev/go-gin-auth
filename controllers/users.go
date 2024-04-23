package controllers

import (
	"net/http"

	"github.com/PontakornDev/ginAuth/models"
	"github.com/PontakornDev/ginAuth/utils"
	"github.com/gin-gonic/gin"
)

func UsersEndPoint(router *gin.RouterGroup) {
	route := router.Group("/users")
	route.POST("/register", register)
	// route.GET("/getProduct/:id", GetProduct)
	// route.GET("/getAllProduct", GetAllProduct)
	// route.PUT("/updateProduct", UpdateProduct)
	// route.DELETE("/deleteProduct/:id", DeleteProduct)
}

func register(ctx *gin.Context) {
	var req = &models.Users{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessMessage(utils.DataObject{
		Title: "Register success",
	}))

}
