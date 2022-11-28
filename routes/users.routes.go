package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JeanVittory/api-rest/db"
	"github.com/JeanVittory/api-rest/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	var usersFromDB = db.DB.Find(&users)
	var error = usersFromDB.Error

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
	}

	json.NewEncoder(w).Encode(&users)

	w.Write([]byte("GetUsersHandler"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var user models.Users
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var user models.Users
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Delete(&user)

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)

	var createdUser = db.DB.Create(&user)
	var error = createdUser.Error

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUserHandler"))
}
