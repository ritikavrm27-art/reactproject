package routes

import (
	"net/http"
	"blogapi/controllers"
	//"github.com/gorilla/mux"
)

func RegisterRoutes() {
	http.HandleFunc("/blogs", controllers.GetBlogs)
	http.HandleFunc("/blog/create", controllers.CreateBlog)
	http.HandleFunc("/blog/delete", controllers.DeleteBlog)
}