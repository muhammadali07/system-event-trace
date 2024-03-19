package app

import (
	"github.com/muhammadali07/service-grap-go-api/services/acc/repository"
	"github.com/sirupsen/logrus"
)

type AccountApp struct {
	repo *repository.Accountepository
	log  *logrus.Logger
}

func InitApp(repo *repository.Accountepository, log *logrus.Logger) *AccountApp {
	return &AccountApp{
		repo: repo,
		log:  log,
	}
}
