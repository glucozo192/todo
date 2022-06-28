package routes

import (
	"TOGO/controllers"
	"TOGO/middleware"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user", controllers.Signup()).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.EditAUser()).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	//----------------------------------------------------------------
	router.Handle("/users", controllers.GetAllUser()).Methods("GET")
	router.HandleFunc("/login", controllers.Login()).Methods("POST")
	router.Handle("/getme", middleware.AuthMiddleware(controllers.GetMe())).Methods("GET")
}
