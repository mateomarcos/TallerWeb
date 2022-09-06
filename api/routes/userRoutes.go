package routes

import (
	controller "Portfolio/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/", controller.GetUser)
	//incomingRoutes.GET("/users/:user_id", controller.GetExtUser()) De momento no, solo vamos a ver los proyectos de otros usuarios y no su info

}
