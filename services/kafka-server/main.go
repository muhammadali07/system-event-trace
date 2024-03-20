package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

// KafkaServer adalah struktur untuk mengelola koneksi ke Kafka server.
type KafkaServer struct {
	reader *kafka.Reader
}

// NewKafkaServer membuat instance baru dari KafkaServer.
func NewKafkaServer(topic string) *KafkaServer {
	// Konfigurasi pembaca Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "kafka-server-group", // Sesuaikan dengan kebutuhan Anda
		Topic:   topic,
	})

	return &KafkaServer{reader: reader}
}

// Close menutup koneksi ke Kafka server.
func (s *KafkaServer) Close() {
	s.reader.Close()
}

// ConsumeMessage mengonsumsi pesan dari Kafka server.
func (s *KafkaServer) ConsumeMessage(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			message, err := s.reader.FetchMessage(ctx)
			if err != nil {
				log.Println("Error fetching message:", err)
				continue
			}
			log.Printf("Received message: %s\n", message.Value)
			// Handle message, misalnya dengan melemparkan ke Fiber untuk diproses
		}
	}
}

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Inisialisasi Kafka server
	kafkaServer := NewKafkaServer("test_topic")
	defer kafkaServer.Close()

	// Buat context untuk Kafka server
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Mulai goroutine untuk mengonsumsi pesan dari Kafka server
	go kafkaServer.ConsumeMessage(ctx)

	// Endpoint untuk mengirim pesan ke Kafka server
	app.Post("/produce", func(c *fiber.Ctx) error {
		message := c.Body()
		fmt.Println(message)
		// Kirim pesan ke Kafka server
		// TODO: Tambahkan logika untuk mengirim pesan ke Kafka
		return c.SendString("Message sent to Kafka server")
	})

	// Menjalankan server Fiber
	go func() {
		if err := app.Listen(":4000"); err != nil {
			log.Fatalf("Failed to start Fiber server: %v", err)
		}
	}()

	// Menangani sinyal shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down...")
}
