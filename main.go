package main

import (
	"go-upload-excelize/database"
	"go-upload-excelize/handlers"
	"log"
)

func main() {
	if database.CheckingConnection() == false {
		log.Fatal("Missing database connection")

		return
	}

	handlers.MainHandlers()
}
