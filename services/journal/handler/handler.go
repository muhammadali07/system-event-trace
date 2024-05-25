package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/muhammadali07/system-event-trace/services/journal/app"
	"github.com/muhammadali07/system-event-trace/services/journal/pkg/log"
	"github.com/muhammadali07/system-event-trace/services/journal/pkg/utils"
	"github.com/sirupsen/logrus"
)

type HandlerKafka struct {
	app       app.JournalServicePort
	log       *log.Logger
	validator *validator.Validate
	address   string
}

func (h *HandlerKafka) Start() error {
	cfg, err := utils.InitConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"config": cfg,
		}).Warn(err.Error())
	}
	err = h.RunServiceJournal()
	if err != nil {
		h.log.Error(logrus.Fields{"err": err.Error()}, nil, "Error RunServiceJournal")
		return fmt.Errorf(err.Error())
	}

	return nil
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
