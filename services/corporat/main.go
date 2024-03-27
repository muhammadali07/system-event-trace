package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/muhammadali07/service-grap-go-api/services/corporat/database"
	"github.com/muhammadali07/service-grap-go-api/services/corporat/graphql"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	// Initialize Fiber
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// GraphQL endpoint
	app.Post("/graphql", graphql.GraphqlHandler(database.DB))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port
	}
	log.Fatal(app.Listen(":" + port))
}
