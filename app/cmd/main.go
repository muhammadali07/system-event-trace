package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	// Koneksi database
	db, err := sql.Open("postgres", "postgres://localhost:5432/mydb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Repository user
	userRepository := NewUserRepository(db)

	// Repository post
	postRepository := NewPostRepository(db)

	// Skema GraphQL
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType(userRepository, postRepository),
		Mutation: MutationType(userRepository, postRepository),
	})
	if err != nil {
		panic(err)
	}

	// Handler GraphQL
	h := handler.New(&handler.Config{
		Schema: &schema,
	})

	// Server HTTP
	http.HandleFunc("/graphql", h.ServeHTTP)
	fmt.Println("Menjalankan di port 8080")
	http.ListenAndServe(":8080", nil)
}
