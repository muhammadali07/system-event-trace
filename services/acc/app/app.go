package app

import (
	"github.com/muhammadali07/system-event-trace/services/acc/repository"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

type AccountApp struct {
	repo         *repository.Accountepository
	log          *logrus.Logger
	tracer       trace.Tracer
	traceEnabled bool
}

func InitApp(repo *repository.Accountepository, log *logrus.Logger, tracer trace.Tracer, traceEnabled bool) *AccountApp {
	return &AccountApp{
		repo:         repo,
		log:          log,
		tracer:       tracer,
		traceEnabled: traceEnabled,
	}
}
