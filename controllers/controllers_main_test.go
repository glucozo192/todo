package controllers_test

import (
	"TOGO/configs"
	"TOGO/routes"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {

	router := mux.NewRouter()
	code := m.Run()
	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)
	routes.TaskRoute(router)

	defer os.Exit(code)
}
