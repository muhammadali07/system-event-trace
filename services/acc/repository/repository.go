package repository

import (
	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GLRepository struct {
	log *logrus.Logger
	db  *gorm.DB
}

func migrateTransaksi(db *gorm.DB) {
	db.AutoMigrate(models.Transaksi{})
}

func InitRepository(db *gorm.DB, log *logrus.Logger) *GLRepository {
	migrateTransaksi(db)
	return &GLRepository{
		db:  db,
		log: log,
	}
}
