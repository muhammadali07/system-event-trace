package app

import "github.com/muhammadali07/service-grap-go-api/services/journal/pkg/log"

type JournalApplication struct {
	datastore JournalDatastorePort
	log       *log.Logger
}

func InitApplication(datastore JournalDatastorePort, log *log.Logger) *JournalApplication {
	return &JournalApplication{
		datastore: datastore,
		log:       log,
	}
}
