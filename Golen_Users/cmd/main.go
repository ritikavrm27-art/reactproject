package main

import (
	"log"
	"net/http"
	"userapi/config"
	"userapi/routes"
)

func main() {

	config.ConnectDB()

	routes.RegisterRoutes()

	log.Println("Server started on port 8002")

	http.ListenAndServe(":8002", nil)
}