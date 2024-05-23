package main

import (
	"context"
	"net/http"
	"os"

	telemetry "github.com/muhammadali07/system-event-trace/build/opentelemetry"
	"github.com/muhammadali07/system-event-trace/services/journal/app"
	"github.com/muhammadali07/system-event-trace/services/journal/datastore"
	"github.com/muhammadali07/system-event-trace/services/journal/handler"
	"github.com/muhammadali07/system-event-trace/services/journal/pkg/log"
	"github.com/muhammadali07/system-event-trace/services/journal/pkg/utils"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer            trace.Tracer
	propagator        propagation.TextMapPropagator
	statusOptions     = []int{http.StatusOK, http.StatusBadRequest}
	telemetryEndpoint = os.Getenv("TELEMETRY_ENDPOINT")
)

func main() {
	cfg, err := utils.InitConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"config": cfg,
		}).Warn(err.Error())
	}

	if telemetryEndpoint == "" {
		telemetryEndpoint = ""
	}

	ctx := context.Background()
	propagator = telemetry.NewTelemetryPropagators()
	tp := telemetry.NewHTTPTelemetryProvider(telemetryEndpoint, "gl-service", ctx)
	tracer = tp.Tracer("server")

	utils.ConnectDB(cfg.DatabaseDriver, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort, cfg.Database)

	logger := log.NewLogger(cfg.KafkaServiceName)
	journalds := datastore.InitDatastore(logger, tracer, true)
	journalapp := app.InitApplication(journalds, logger, tracer, true)
	consumerService := handler.InitHandlerKafka(logger, journalapp, tracer, true)

	// dsn := datastore.InitDatastore(cfg.DatabaseDriver, cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePassword, cfg.Database, cfg.DatabasePort, cfg.DatabaseSchema, logger)
	// app := app.InitApplication(dsn, logger)
	// consumerService := handler.InitHandlerKafka(cfg.KafkaHost, cfg.KafkaPort, app, logger)

	// running service consumer
	consumerService.Start()
}
