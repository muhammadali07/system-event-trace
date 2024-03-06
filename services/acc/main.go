package main

import (
	"context"
	"fmt"

	"services/acc/internal/app"
)

func main() {
	ctx := context.Background()

	// Konfigurasi awal aplikasi
	cfg, err := app.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Memuat dependensi
	db, err := app.NewDB(cfg.DB)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Menjalankan fungsi utama aplikasi
	app.Run(ctx, cfg, db)
}
