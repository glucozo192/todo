package controllers_test

import (
	"TOGO/configs"
	"TOGO/controllers"
	"TOGO/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

var a App

var tokenMain string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjA4NzgxNzMsImlkIjoiNjJiYWJkNTA0OTBjMmE0ODc4MTViYzcxIn0.wcvPD8ly0YMoSiRrUkCQ3upS2xjby4hOU7LLybk7pqQ"

func TestMain(m *testing.M) {
	a.Router = mux.NewRouter()
	//run database
	configs.ConnectDB()

	UserRoute((a.Router))
	//routes
	code := m.Run()
	defer os.Exit(code)
}

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/{Id}", controllers.GetUser()).Methods("GET")
	router.HandleFunc("/user/{Id}", controllers.DeleteUser()).Methods("DELETE")
	router.Handle("/task/{id}", middleware.AuthMiddleware(controllers.GetOneTask())).Methods("GET")
	router.Handle("/task/{id}", middleware.AuthMiddleware(controllers.UpdateTask())).Methods("PUT")
	router.Handle("/task/status/{id}", middleware.AuthMiddleware(controllers.UpdateTaskStatus())).Methods("PUT")

}

func ExcuteRoute(r *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.Handler(a.Router)
	handler.ServeHTTP(rr, r)
	return rr
}
