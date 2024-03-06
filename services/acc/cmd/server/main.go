package main

import (
	"context"
	"fmt"

	"services/acc/internal/app/cmd/server"
)

func main() {
	ctx := context.Background()

	// Menjalankan sub-perintah/command server
	if err := server.Run(ctx); err != nil {
		fmt.Println("Error running server:", err)
	}
}
