package repository

import (
	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Accountepository struct {
	log *logrus.Logger
	db  *gorm.DB
}

func migrateTransaksi(db *gorm.DB) {
	db.AutoMigrate(models.Account{})
}

func InitRepository(db *gorm.DB, log *logrus.Logger) *Accountepository {
	migrateTransaksi(db)
	return &Accountepository{
		db:  db,
		log: log,
	}
}
