package main

import (
	"log"
	"net/http"
	"os"

	"DZ1/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                   string
	itemResourcePrefix     string = apiPrefix + "/item"
	manyItemResourcePrefix string = apiPrefix + "/items"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildItemResource(router, itemResourcePrefix)
	utils.BuildManyItemsResourcePrefix(router, manyItemResourcePrefix)

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
