package routes

import (
	controller "Portfolio/controllers"

	"github.com/gin-gonic/gin"
)

/*
Input: Group of routes of the gin router/engine.
Output: HTTP methods and routes added to our router, which specific function to call when they are called.
*/
func ProjectRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/projects", controller.GetProjects)
	incomingRoutes.POST("/projects", controller.NewProject)
	incomingRoutes.GET("/:username/projects", controller.GetExtProjects) // ":" before a name specifies a parameter for the route.
	incomingRoutes.DELETE("/projects/:name", controller.DeleteProject)   // ":" before a name specifies a parameter for the route.
}
