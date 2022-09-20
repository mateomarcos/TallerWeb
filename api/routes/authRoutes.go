package routes

import (
	controller "Portfolio/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", controller.Signup)
	incomingRoutes.POST("/login", controller.Login)
	//incomingRoutes.GET("/logout", controller.Logout)
}
