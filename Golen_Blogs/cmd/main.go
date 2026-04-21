package main

import (
	"log"
	"net/http"
	"blogapi/config"
	"blogapi/routes"
)

func main() {

	config.ConnectDB()

	routes.RegisterRoutes()

	log.Println("Server started on port 8001")

	http.ListenAndServe(":8001", nil)
}