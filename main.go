package main

import (
	"net/http"

	"github.com/JeanVittory/api-rest/db"
	"github.com/JeanVittory/api-rest/models"
	"github.com/JeanVittory/api-rest/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.Users{})
	var router = mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/task", routes.GetTasksHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/user", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", routes.UpdateUserHandler).Methods("PUT")
	http.ListenAndServe(":8080", router)
}
