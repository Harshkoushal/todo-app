package services

import (
	"context"
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
