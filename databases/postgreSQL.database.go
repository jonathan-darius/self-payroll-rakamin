package databases

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"self-payroll/models"
)

var (
	DBClient *gorm.DB = dbInit()
)

func dbInit() *gorm.DB {
	err := godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println(dbURL)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cant Connect To Databases: ", err)
	}
	db.AutoMigrate(&models.Position{})
	return db
}
