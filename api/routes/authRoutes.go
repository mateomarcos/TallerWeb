package routes

import (
	controller "Portfolio/controllers"

	"github.com/gin-gonic/gin"
)

/*
Input: Group of routes of the gin router/engine.
Output: HTTP methods and routes added to our router, which specific function to call when they are called.
*/
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", controller.Signup)
	incomingRoutes.POST("/login", controller.Login)
}
