package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/muhammadali07/system-event-trace/services/journal/pkg/utils"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func (h *HandlerKafka) RunServiceJournal() (err error) {
	cfg, err := utils.InitConfig()
	if err != nil {
		h.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, "failed to open config")
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
	logrus.Info(m)
	for topic := range m {
		handler, _ := h.RouteTopic(topic, nil) // Get the handler function
		if handler != nil {                    // Only proceed if a handler is found
			logrus.Info("topic_from_kafka: ", topic)
			// go h.consumeMessages(topic, brokers) // Pass the handler to consumeMessages
			res, err := h.consumeMessages(topic, brokers)
			if err != nil {
				panic(err.Error())
			}
			h.log.Info(logrus.Fields{
				"resutl": res,
			}, nil, "result handler by topic")
		}
	}
	// 	go consumeMessages(topic, brokers) -> manually getting 1 by 1 topic by hardcode

	// select {}
	return
}

func (h *HandlerKafka) consumeMessages(topic string, brokers []string) (response any, err error) {
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
		response, err = h.RouteTopic(topic, message.Value)
		if err != nil {
			h.log.Info(logrus.Fields{
				"response": response,
				"err":      err,
			}, nil, "handling route topic")
		}
		// fmt.Println(response)
	}
}
