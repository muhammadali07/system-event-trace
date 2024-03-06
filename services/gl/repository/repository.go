package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ihsansolusi.co.id/information-centre/backend/models"
)

type ICRepository struct {
	log *logrus.Logger
	db  *gorm.DB
}

func migrateEmployee(db *gorm.DB) {
	db.AutoMigrate(models.Employee{}, models.Address{}, models.Skills{})
}

func migrateProject(db *gorm.DB) {
	db.AutoMigrate(models.Client{}, models.Project{}, models.Delivery{}, models.Engineer{})
}

func migrateActivity(db *gorm.DB) {
	db.AutoMigrate(models.Activity{})
}

func InitRepository(db *gorm.DB, log *logrus.Logger) *ICRepository {
	migrateEmployee(db)
	migrateProject(db)
	migrateActivity(db)
	return &ICRepository{
		db:  db,
		log: log,
	}
}
