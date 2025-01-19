package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title" bson:"title"`
	Priority string             `json:"priority" bson:"priority"`
	Time     string             `json:"time" bson:"time"`
	Status   bool               `json:"status" bson:"status"`
}

