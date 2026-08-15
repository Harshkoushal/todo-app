package services

import (
	"context"
	"errors"
	"time"

	"todo-app/database"
	"todo-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(userID primitive.ObjectID, req models.CreateTodoRequest) (*models.Todo, error) {
	collection := database.GetCollection("todos")

	todo := models.Todo{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	result, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}

	todo.ID = result.InsertedID.(primitive.ObjectID)
	return &todo, nil
}

func GetTodos(userID primitive.ObjectID) ([]models.Todo, error) {
	collection := database.GetCollection("todos")

	cursor, err := collection.Find(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}

	todos := []models.Todo{}
	err = cursor.All(context.Background(), &todos)
	return todos, err
}
func GetTodoByID(userID, todoID primitive.ObjectID) (*models.Todo, error) {
	collection := database.GetCollection("todos")

	var todo models.Todo
	err := collection.FindOne(context.Background(), bson.M{"_id": todoID, "userId": userID}).Decode(&todo)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	return &todo, nil
}

func UpdateTodo(userID, todoID primitive.ObjectID, req models.CreateTodoRequest) (*models.Todo, error) {
	collection := database.GetCollection("todos")

	update := bson.M{
		"title":       req.Title,
		"description": req.Description,
	}

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": todoID, "userId": userID},
		bson.M{"$set": update},
	)
	if err != nil {
		return nil, err
	}

	return GetTodoByID(userID, todoID)
}

func DeleteTodo(userID, todoID primitive.ObjectID) error {
	collection := database.GetCollection("todos")

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": todoID, "userId": userID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("todo not found")
	}

	return nil
}
