package router

import (
	"net/http"

	"github.com/PontakornDev/ginAuth/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello demo auth api")
	})

	return router
}

func SetUpRouteGroup(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	controllers.UsersEndPoint(v1)
	controllers.AuthEndPoint(v1)
}
