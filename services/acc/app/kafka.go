package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/muhammadali07/service-grap-go-api/services/acc/pkg/utils"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func (a *AccountApp) SendMessageToKafka(params models.ReqSendingKafka) (err error) {
	cfg, _ := utils.InitConfig()
	broker := fmt.Sprintf("%v:%v", cfg.KafkaHost, cfg.KafkaPort)

	// register new dinamic topic
	err = a.CreateNewTopicKafka(params.Topic, broker)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}
	// Set up Kafka writer
	topic := params.Topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	// Create message payload
	payload, err := json.Marshal(params.Data)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}

	// Write message to Kafka
	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	})
	if err != nil {
		err = fmt.Errorf(err.Error())
		return
	}
	a.log.WithFields(logrus.Fields{
		"payload": string(payload),
	}).Info("Message sent to kafka")

	return
}

func (a *AccountApp) CreateNewTopicKafka(topic string, broker string) (err error) {
	a.log.WithFields(logrus.Fields{
		"topic":  topic,
		"broker": broker,
	}).Info("processing to create new topic")

	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"topic":  topic,
			"broker": broker,
			"err":    err.Error(),
		})
		return
	}

	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"topic":  topic,
			"broker": broker,
			"err":    err.Error(),
		}).Error("Failed to create new topic")
	} else {
		a.log.WithFields(logrus.Fields{
			"topic": topic,
		}).Info("successfullty to create new topic")
	}
	return
}
