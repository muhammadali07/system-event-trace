package app

import (
	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/log"
	"go.opentelemetry.io/otel/trace"
)

type JournalApplication struct {
	datastore    JournalDatastorePort
	log          *log.Logger
	tracer       trace.Tracer
	traceEnabled bool
}

func InitApplication(datastore JournalDatastorePort, log *log.Logger, tracer trace.Tracer, traceEnabled bool) *JournalApplication {
	return &JournalApplication{
		datastore:    datastore,
		log:          log,
		tracer:       tracer,
		traceEnabled: traceEnabled,
	}
}
