package datastore

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"

	"github.com/muhammadali07/service-grap-go-api/services/journal/pkg/log"
	"github.com/muhammadali07/system-event-trace/services/journal/pkg/utils"
)

type JournalDatabase struct {
	db     *gorm.DB
	log    *log.Logger
	tracer trace.Tracer
}

func (f *JournalDatabase) Begin() (tx *gorm.DB, err error) {
	tx = f.db.Begin()
	if tx.Error != nil {
		remark := "failed to start transaction"
		f.log.Error(logrus.Fields{
			"error": tx.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
	}
	return
}

func (f *JournalDatabase) Rollback(tx *gorm.DB) {
	err := tx.Rollback()
	if err != nil {
		f.log.Error(logrus.Fields{
			"error": err.Error,
		}, nil, "failed to rollback transaction")
	}
}

func (f *JournalDatabase) Commit(tx *gorm.DB) {
	err := tx.Commit()
	if err != nil {
		f.log.Error(logrus.Fields{
			"error": err.Error,
		}, nil, "failed to commit transaction")
	}
}

func InitDatastore(log *log.Logger, tracer trace.Tracer, traceEnabled bool) *JournalDatabase {

	return &JournalDatabase{
		db:     utils.DBInstance,
		log:    log,
		tracer: tracer,
	}
}
