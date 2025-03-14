package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {

	var err error
	var env = os.Getenv("ENV")
	var logLevel = int(logger.Silent)

	dataSourceName := fmt.Sprintf("host='%s' user='%s' password='%s' dbname='%s' port=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"))

	if env == "LOCAL" {
		logLevel = int(logger.Info)
	}

	DB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
	})
	if err != nil {
		log.Fatal("Failed to connect to the Database" + err.Error())
	}
	log.Println("? Connected Successfully to the Database")

}
