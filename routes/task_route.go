package routes

import (
	"TOGO/controllers"
	"TOGO/middleware"

	"github.com/gorilla/mux"
)

func TaskRoute(router *mux.Router) {
	router.Handle("/tasks", controllers.GetAllTask()).Methods("GET")
	router.HandleFunc("/task/{id}", controllers.GetOneTask()).Methods("GET")
	router.Handle("/task", middleware.AuthMiddleware(controllers.CreateTask())).Methods("POST")
	router.Handle("/user-tasks", middleware.AuthMiddleware(controllers.GetTask())).Methods("GET")
	router.Handle("/task/{id}", middleware.AuthMiddleware(controllers.DeleteTask())).Methods("DELETE")
	router.HandleFunc("/task/{id}", controllers.UpdateTask()).Methods("PUT")
	router.Handle("/task/status/{id}", middleware.AuthMiddleware(controllers.UpdateTaskStatus())).Methods("PUT")
}
