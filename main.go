package main

import (
	"TOGO/configs"
	"TOGO/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)
	routes.TaskRoute(router)

	log.Fatal(http.ListenAndServe(":9099", router))
}
