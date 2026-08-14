package main

import (
	"log"
	"net/http"

	"todo-app/database"
	"todo-app/handlers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/auth/register", handlers.Register)
	mux.HandleFunc("POST /api/auth/login", handlers.Login)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
