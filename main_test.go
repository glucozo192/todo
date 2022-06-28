package main

import (
	"TOGO/configs"
	"TOGO/routes"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()
	routes.UserRoute(router)
	routes.TaskRoute(router)
	code := m.Run()
	os.Exit(code)

}
