package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"DZ1/models"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get infos about all books in database")
	if models.IsDBEmpty() {
		writer.WriteHeader(403)
		msg := models.JSONError{Error: "No one items found in store back"}
		err := json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
	} else {
		writer.WriteHeader(200)
		err := json.NewEncoder(writer).Encode(models.DB)
		if err != nil {
			log.Fatal("Can't encode DB, something extremely wrong")
		}
	}
}
