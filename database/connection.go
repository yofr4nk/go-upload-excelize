package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/yofr4nk/file-upload/models"
	"log"
	"os"
)

var DBConnect = initDB()

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
	db, dbErr := gorm.Open("mysql", dataSourceName)

	if dbErr != nil {
		fmt.Println(dbErr)
		panic("failed to connect database")
	}

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&models.Person{})

	fmt.Println("Connection successful")

	return db
}

// CheckingConnection check if the database connection is available
func CheckingConnection() bool {
	checkError := DBConnect.DB().Ping()

	if checkError != nil {
		return false
	}

	return true
}
