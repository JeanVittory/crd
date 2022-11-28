package routes

import "net/http"

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Task route"))
}
