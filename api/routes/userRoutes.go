package routes

import (
	controller "Portfolio/controllers"

	"github.com/gin-gonic/gin"
)

/*
Input: Group of routes of the gin router/engine.
Output: HTTP methods and routes added to our router, which specific function to call when they are called.

This is not actually used in the frontend.
*/
func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/", controller.GetUser)
}
