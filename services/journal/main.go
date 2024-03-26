package main

import (
	"github.com/muhammadali07/service-grap-go-api/services/journal/app"
	"github.com/muhammadali07/service-grap-go-api/services/journal/datastore"
	"github.com/muhammadali07/service-grap-go-api/services/journal/handler"
	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/log"
	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := utils.InitConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"config": cfg,
		}).Warn(err.Error())
	}

	logger := log.NewLogger(cfg.KafkaServiceName)
	dsn := datastore.InitDatastore(cfg.DatabaseDriver, cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePassword, cfg.Database, cfg.DatabasePort, cfg.DatabaseSchema, logger)
	app := app.InitApplication(dsn, logger)
	consumerService := handler.InitHandlerKafka(cfg.KafkaHost, cfg.KafkaPort, app, logger)

	// running service consumer
	consumerService.Start()
}
