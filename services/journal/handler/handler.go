package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/muhammadali07/service-grap-go-api/services/journal/app"
	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/log"
	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/utils"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type HandlerKafka struct {
	app       app.JournalServicePort
	log       *log.Logger
	validator *validator.Validate
	address   string
}

func (h *HandlerKafka) Start() {
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

}

func InitHandlerKafka(host string, port int, app app.JournalServicePort, log *log.Logger) *HandlerKafka {
	address := fmt.Sprintf("%s:%d", host, port)
	return &HandlerKafka{
		address:   address,
		log:       log,
		app:       app,
		validator: validator.New(),
	}
}
