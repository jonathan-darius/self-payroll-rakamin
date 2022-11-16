package routes

import (
	"github.com/gin-gonic/gin"
	"self-payroll/controllers"
)

func EmployeesRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/", controllers.CreateEmployees())
	incomingRoutes.DELETE("/:id", controllers.DeleteEmployees())
	incomingRoutes.GET("/", controllers.AllEmployees())
	incomingRoutes.GET("/:id", controllers.DetailEmployees())
	incomingRoutes.POST("/withdraw", controllers.WithdrawBalance())
}
