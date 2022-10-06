package routes

import (
	controller "Portfolio/controllers"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/projects", controller.GetProjects)
	incomingRoutes.POST("/projects", controller.NewProject)
	incomingRoutes.GET("/:username/projects", controller.GetExtProjects)
	incomingRoutes.DELETE("/projects/:name", controller.DeleteProject)
}
