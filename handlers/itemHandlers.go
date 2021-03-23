package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"DZ1/models"
	"github.com/gorilla/mux"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.JSONError{Error: "do not use parameter ID as uncasted to int type"}
		err := json.NewEncoder(writer).Encode(msg); if err != nil {
			log.Fatal("Can't encode DB, something extremely wrong")
		}
		return
	}

	item, ok := models.FindItemById(id)
	log.Println("Get item with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.JSONError{Error: "Item with that id not found"}
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
	} else {
		writer.WriteHeader(200)
		err = json.NewEncoder(writer).Encode(item)
		if err != nil {
			log.Fatal("Can't encode item, something extremely wrong")
		}
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new item ....")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.JSONError{Error: "do not use parameter ID as uncasted to int type"}
		err := json.NewEncoder(writer).Encode(msg); if err != nil {
			log.Fatal("Can't encode DB, something extremely wrong")
		}
		return
	}
	var item models.Item

	err = json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		msg := models.JSONError{Error: "provideed json file is invalid"}
		writer.WriteHeader(400)
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
		return
	}
	if models.AddItemToDB(id, item) {
		writer.WriteHeader(201)
		err = json.NewEncoder(writer).Encode(item)
		if err != nil {
			log.Fatal("Can't encode item, something extremely wrong")
		}
	} else {
		writer.WriteHeader(400)
		msg := models.JSONError{Error: "Item with that id already exists"}
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
	}
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.JSONError{Error: "do not use parameter ID as uncasted to int type"}
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
		return
	}
	var newItem models.Item
	err = json.NewDecoder(request.Body).Decode(&newItem)
	if err != nil {
		msg := models.JSONError{Error: "provided json file is invalid"}
		writer.WriteHeader(400)
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
		return
	}
	if models.UpdateItemInDB(id, newItem) {
		writer.WriteHeader(202)
		err = json.NewEncoder(writer).Encode(newItem)
		if err != nil {
			log.Fatal("Can't encode item, something extremely wrong")
		}
	} else {
		log.Println("item not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.JSONError{Error: "item with that ID not found"}
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
		return
	}
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Deleting item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.JSONError{Error: "do not use parameter ID as uncasted to int type"}
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
		return
	}
	ok := models.DeleteItemFromDB(id)
	if !ok {
		log.Println("item not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.JSONError{Error: "item with that ID not found"}
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("Can't encode error, something extremely wrong")
		}
		return
	}
	msg := models.JSONError{Error: "Item deleted"}
	err = json.NewEncoder(writer).Encode(msg)
	if err != nil {
		log.Fatal("Can't encode error, something extremely wrong")
	}
}
