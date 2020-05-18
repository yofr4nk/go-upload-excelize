package routers

import (
	"encoding/json"
	"github.com/yofr4nk/file-upload/files"
	"net/http"
)

// UploadFile route to process file before save it
func UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	file, _, err := r.FormFile("people_file")

	if err != nil {
		http.Error(w, "Error Retrieving the File: "+err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	people, processErr := files.ProcessPeople(file)

	if processErr != nil {
		http.Error(w, "Something went wrong: "+processErr.Error(), http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}
