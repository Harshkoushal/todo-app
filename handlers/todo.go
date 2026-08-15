package handlers

import (
	"encoding/json"
	"net/http"

	"todo-app/models"
	"todo-app/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	var req models.CreateTodoRequest
	json.NewDecoder(r.Body).Decode(&req)

	todo, err := services.CreateTodo(userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	todos, err := services.GetTodos(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
func GetTodo(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	todoID, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	todo, err := services.GetTodoByID(userID, todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	todoID, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	var req models.CreateTodoRequest
	json.NewDecoder(r.Body).Decode(&req)

	todo, err := services.UpdateTodo(userID, todoID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Context().Value("userID").(string)
	userID, _ := primitive.ObjectIDFromHex(userIDStr)

	todoID, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid todo id", http.StatusBadRequest)
		return
	}

	if err := services.DeleteTodo(userID, todoID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "todo deleted"})
}
