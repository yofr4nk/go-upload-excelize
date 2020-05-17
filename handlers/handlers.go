package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yofr4nk/file-upload/routers"
	"log"
	"net/http"
	"os"
)

// MainManagement set the main config for routers
func MainHandlers() {
	router := mux.NewRouter()

	router.HandleFunc("/upload-file", routers.UploadFile).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
