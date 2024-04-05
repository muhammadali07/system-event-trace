// package main

// import (
// 	"database/sql"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	"github.com/gofiber/fiber"
// )

// func main() {

// 	// Buat GraphQL server
// 	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{DB: db}}))

// 	// Buat Fiber app
// 	app := fiber.New()

// 	// Playground
// 	app.Get("/playground", playground.Handler("GraphQL Playground", "/query"))

// 	// GraphQL endpoint
// 	app.Post("/query", srv.ServeHTTP)

// 	// Listen
// 	app.Listen(":8080")
// }

// type Resolver struct {
// 	DB *sql.DB
// }

// func (r *Resolver) Users() ([]User, error) {
// 	// Query database
// 	rows, err := r.DB.Query("SELECT * FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Scan rows
// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.ID, &user.Name, &user.Email)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

// func (r *Resolver) CreateUser(name, email string) (User, error) {
// 	// Insert ke database
// 	stmt, err := r.DB.Prepare("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email")
// 	if err != nil {
// 		return User{}, err
// 	}
// 	defer stmt.Close()

// 	// Jalankan query
// 	row := stmt.QueryRow(name, email)

// 	// Scan hasil
// 	var user User
// 	err = row.Scan(&user.ID, &user.Name, &user.Email)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }
