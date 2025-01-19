package services

import (
	"context"
	"errors"
	

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"todo-app/config"
	"todo-app/models"

)


func CreateTodo(todo models.Todo)(primitive.ObjectID, error){
	result , err := config.TodoCollection.InsertOne(context.Background(), todo)

	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func GetTodoByDate()([]models.Todo, error){
	var todos []models.Todo
	cursor, err := config.TodoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()){
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil{
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}


func UpdateTodo(id primitive.ObjectID, updateData models.Todo) error {
	// Exclude _id from the updateData
	update := bson.M{
		"$set": bson.M{
			"status":   updateData.Status,
		},
	}

	// Update operation
	result, err := config.TodoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id}, // Filter by ID
		update,            // Update query
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no document found with the given id")
	}
	return nil
}


func DeleteTodo(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	result, err := config.TodoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no document found with the given id")
	}
	return nil
}