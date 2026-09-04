package main

import (
	"log"
	"net/http"

	"todo-app/database"
	"todo-app/handlers"
	"todo-app/middleware"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/auth/register", handlers.Register)
	mux.HandleFunc("POST /api/auth/login", handlers.Login)

	mux.Handle("POST /api/todos", middleware.Authenticate(http.HandlerFunc(handlers.CreateTodo)))
	mux.Handle("GET /api/todos", middleware.Authenticate(http.HandlerFunc(handlers.GetTodos)))
	mux.Handle("GET /api/todos/{id}", middleware.Authenticate(http.HandlerFunc(handlers.GetTodo)))
	mux.Handle("PUT /api/todos/{id}", middleware.Authenticate(http.HandlerFunc(handlers.UpdateTodo)))
	mux.Handle("DELETE /api/todos/{id}", middleware.Authenticate(http.HandlerFunc(handlers.DeleteTodo)))

	handler := middleware.EnableCORS(mux)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", handler)
}
