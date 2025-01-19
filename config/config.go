package config

import (
	 "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "context"
	"os"
 	"github.com/joho/godotenv"
)

var Client *mongo.Client
var TodoCollection *mongo.Collection

func ConnectToMongoDB(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	clinetOption := options.Client().ApplyURI(databaseURL)
	client, err := mongo.Connect(context.Background(), clinetOption)

	if err != nil {
		log.Fatal(err)
	}
	Client = client

	TodoCollection = client.Database("todo_db").Collection("todos_native")

	err = Client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect to MongoDB!")
}