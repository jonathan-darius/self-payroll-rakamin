package routes

import (
	"github.com/gin-gonic/gin"
	"self-payroll/controllers"
)

func PositionRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/", controllers.CreatePosition())
	incomingRoutes.DELETE("/:id", controllers.DeletePosition())
	incomingRoutes.GET("/:id", controllers.SearchPosition())
	incomingRoutes.GET("/", controllers.AllPosition())
}
