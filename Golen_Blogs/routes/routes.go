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
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/upload", controllers.UploadFile)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/app/uploads"))))
}