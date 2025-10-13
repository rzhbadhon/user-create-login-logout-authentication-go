package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine){
	router.POST("/signup", SignUp())
	router.POST("/login", Login())
}