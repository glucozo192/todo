package routes

import (
	"TOGO/controllers"
	"TOGO/middleware"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user/{Id}", controllers.GetUser()).Methods("GET")
	router.Handle("/user", middleware.AuthMiddleware(controllers.UpdateMe())).Methods("PUT")
	router.HandleFunc("/user/{Id}", controllers.DeleteUser()).Methods("DELETE")
	//----------------------------------------------------------------
	router.HandleFunc("/user/signup", controllers.Signup()).Methods("POST")
	router.Handle("/users", controllers.GetAllUser()).Methods("GET")
	router.HandleFunc("/user/login", controllers.Login()).Methods("POST")
	router.Handle("/me", middleware.AuthMiddleware(controllers.GetMe())).Methods("GET")
}
