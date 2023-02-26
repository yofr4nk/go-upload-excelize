package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-upload-excelize/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DBConnection = initDB()

func initDB() *gorm.DB {
	// Loading Env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables", err.Error())
	}

	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dataSourceName := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbName + "?parseTime=True"
	db, dbErr := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if dbErr != nil {
		fmt.Println(dbErr)
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Person{})

	fmt.Println("Connection successful")

	return db
}

// CheckingConnection check if the database connection is available
func CheckingConnection() bool {
	_, err := DBConnection.DB()

	if err != nil {
		return false
	}

	return true
}
