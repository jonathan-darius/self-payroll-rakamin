package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"self-payroll/routes"
)

var (
	router *gin.Engine
	port   string
)

func init() {
	err := godotenv.Load(".env")

	port = os.Getenv("PORT")
	fmt.Println(port)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if port == "" {
		fmt.Println("Use 8000")
		port = "8000"
	}
	router = gin.Default()

	//router.Use(gin.Logger())
}

func main() {
	position := router.Group("/position")
	company := router.Group("/company")
	employees := router.Group("/employee")
	transactions := router.Group("/transactions")

	routes.PositionRoutes(position)
	routes.CompanyRoutes(company)
	routes.EmployeesRoutes(employees)
	routes.TransactionsRoutes(transactions)
	errRun := router.Run(":" + port)
	if errRun != nil {
		log.Fatal(errRun)
	}
}
