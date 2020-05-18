package main

import (
	"github.com/yofr4nk/file-upload/database"
	"github.com/yofr4nk/file-upload/handlers"
	"log"
)

func main() {
	if database.CheckingConnection() == false {
		log.Fatal("Missing database connection")

		return
	}

	handlers.MainHandlers()
}
