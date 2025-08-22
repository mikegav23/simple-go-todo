package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-go-todo/internal/db"
	"simple-go-todo/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("DATABASE_URL")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close(ctx)

	q := db.New(conn)
	h := handlers.NewHandler(q)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/login", h.LoginHandler)

	r.Get("/users", h.ListUsersHandler)
	r.Post("/users", h.CreateUserHandler)
	r.Put("/users/{userID}", h.UpdateUserHandler)
	r.Delete("/users/{userID}", h.DeleteUserHandler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", r)
}