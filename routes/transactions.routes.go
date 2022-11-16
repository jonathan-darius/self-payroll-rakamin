package routes

import (
	"github.com/gin-gonic/gin"
	"self-payroll/controllers"
)

func TransactionsRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/", controllers.TransactionsGet())
}
