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
	 if port ==""{
		port = ":8000"
	 } 
	config.ConnectToMongoDB()
	router := routes.SetupRouter()
	log.Fatal(http.ListenAndServe( port, corsHandler.Handler(router)))
}