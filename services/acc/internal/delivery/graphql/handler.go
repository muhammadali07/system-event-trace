package main

import (
	"log"
	"os"

	"github.com/muhammadali07/service-grap-go-api/services/acc/internal/delivery/http"
	"github.com/muhammadali07/service-grap-go-api/services/acc/internal/infrastructure/database"
	"github.com/muhammadali07/service-grap-go-api/services/acc/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Koneksi ke database
	db, err := database.NewConnection(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	// Buat use case
	uc := usecase.AccountUseCase(db)

	// Buat server HTTP
	app := fiber.New()

	// Routing HTTP
	http.RegisterHandlers(app, uc)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server berjalan di port", port)
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
