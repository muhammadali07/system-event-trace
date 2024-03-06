package main

import (
    "github.com/muhammadali07/service-grap-go-api/services/kafka/internal/delivery/http"
)

func main() {
    // Initialize and start HTTP server
    http.StartServer()
}
