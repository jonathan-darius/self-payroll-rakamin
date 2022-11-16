package routes

import (
	"github.com/gin-gonic/gin"
	"self-payroll/controllers"
)

func CompanyRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.GET("/", controllers.GetCompany())
	incomingRoutes.POST("/", controllers.CreateOrUpdateCompany())
	incomingRoutes.POST("/topup", controllers.TopUpCompany())
}
