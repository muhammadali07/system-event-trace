package cmd

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	// Import models package (if relevant)

	gql "github.com/muhammadali07/service-grap-go-api/app/gql"
	// Import repository package
)

func main() {
	// Koneksi database
	db, err := sql.Open("postgres", "postgres://localhost:5432/mydb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// postRepository := repository.NewPostRepository(db)

	// Skema GraphQL
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    gql.Query(userRepository, postRepository),
		Mutation: gql.Mutation(userRepository, postRepository),
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
