package routes

import (
	controllers "greenstech/controller"
	accesstoken "greenstech/helper/AccessToken"

	"github.com/gin-gonic/gin"
)

func InitFileHandlerRoutes(router *gin.Engine) {
	route := router.Group("/api/v1/filehandler")
	route.POST("/upload-profile-image", accesstoken.JWTMiddleware(), controllers.PostUploadProfileImage())
	route.POST("/upload-file", accesstoken.JWTMiddleware(), controllers.PostUploadFileController())
}
