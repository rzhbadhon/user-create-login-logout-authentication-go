package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine){
	router.POST("/signup", SignUp())
	router.POST("/login", Login())

	protected := router.Group("/")

	protected.Use(middleware.Authenticate())

	{
		protected.GET("users", controllers.GetUsers())
		protecteD.GET("/user/:id", controllers.GetUser())
	}
}