package routes

import (
	"net/http"
	"userapi/controllers"
	//"github.com/gorilla/mux"
)

func RegisterRoutes() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/user/create", controllers.CreateUser)
	http.HandleFunc("/user/delete", controllers.DeleteUser)
}