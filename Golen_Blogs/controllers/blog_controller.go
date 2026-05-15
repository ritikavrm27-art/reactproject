package controllers

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
	"os"
	"time"

	"blogapi/services"
	"blogapi/models"
)
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	blogs, err := services.GetAllBlogs()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogs)
}
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	var blog models.Blog

	// Decode JSON request body into struct
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call service to insert user
	err = services.InsertBlog(blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog inserted successfully",
	})
}
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog

	// Decode JSON request body into struct
	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.DeleteBlog(blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Blog deleted successfully",
	})

}
func UploadFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.MkdirAll("/app/uploads", os.ModePerm)

	// unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)

	dst, err := os.Create("/app/uploads/" + filename)
	if err != nil {
		http.Error(w, "Save error", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)

	// return URL
	response := map[string]string{
		"url": "http://localhost:30001/files/" + filename,
	}

	json.NewEncoder(w).Encode(response)
}