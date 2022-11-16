package databases

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalln("Cant Connect To Databases: ", err)
	}

	db.AutoMigrate(&models.Position{}, &models.Company{})
	return db
}
