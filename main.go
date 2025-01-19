package main

import (
	"net/http"
	"todo-app/config"
    "todo-app/routes"
    "log"
	"github.com/rs/cors"
    "os"
		"github.com/joho/godotenv"
	)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	corsHandler := cors.Default()
     port := os.Getenv("PORT") 
	config.ConnectToMongoDB()
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(port || ":5000", corsHandler.Handler(router)))
}