package routes

import (
	controllers "greenstech/controller"
	accesstoken "greenstech/helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitSubtrainerRoutes(router *gin.Engine) {
	route := router.Group("/api/v1/subtrainer")
	route.POST("/new", accesstoken.JWTMiddleware(), controllers.NewSubtrainerRegistrationController())
	route.POST("/", accesstoken.JWTMiddleware(), controllers.GetSubtrainerRegistrationController())
}
