package main

import (
	"context"
	"fmt"
	"log"

	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/utils"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// Membuat tipe untuk fungsi yang akan menangani pesan Kafka
type messageHandler func(message kafka.Message)

// Map yang memetakan nama topik ke fungsi yang akan menanganinya
var topicHandlers = map[string]messageHandler{
	"account_topic":     handleAccountMessage,
	"account_get_topic": handleAccountGetMessage,
	"default_topic":     handleDefaultMessage,
}

func main() {
	cfg, err := utils.InitConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"config": cfg,
		}).Warn(err.Error())
	}
	// Brokers represent the Kafka cluster to connect to.
	kafkaAddress := fmt.Sprintf(`%v:%v`, cfg.KafkaHost, cfg.KafkaPort)

	logrus.Info(fmt.Sprintf("Service: %v started successfully 🚀 running on -> %v", cfg.KafkaServiceName, kafkaAddress))

	brokers := []string{kafkaAddress}
	conn, err := kafka.Dial("tcp", fmt.Sprintf(`%v`, brokers[0]))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for topic := range m {
		// Periksa apakah topik ada dalam map topicHandlers sebelum menangani pesan.
		if _, ok := topicHandlers[topic]; ok {
			logrus.Info("topic_from_kafka: ", topic)

			// Konsumsi pesan dalam goroutine agar tidak memblokir loop utama.
			go consumeMessages(topic, brokers)
		}
	}
	// 	go consumeMessages(topic, brokers) -> manually getting 1 buy 1 topic by hardcode

	select {}

}

func consumeMessages(topic string, brokers []string) {
	// Konfigurasi pembaca Kafka dengan topik yang sesuai
	config := kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  fmt.Sprintf("%s-consumer-group", topic),
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}
		log.Printf("Received message from topic %s: %s\n", topic, message.Value)

		// Panggil fungsi penangan pesan berdasarkan topik
		if handler, ok := topicHandlers[topic]; ok {
			handler(message)
		} else {
			handleDefaultMessage(message)
		}
	}
}

func handleAccountMessage(message kafka.Message) {
	log.Println("Handling message from account_topic:", string(message.Value))
	// Tambahkan logika penyimpanan data ke dalam database di sini
}

func handleAccountGetMessage(message kafka.Message) {
	log.Println("Handling message from account_get_topic:", string(message.Value))
	// Tambahkan logika pengambilan data dari database di sini
}

func handleDefaultMessage(message kafka.Message) {
	log.Println("Unknown topic or no handler specified for topic:", message.Topic)
}
