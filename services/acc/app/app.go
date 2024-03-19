package app

import (
	"github.com/muhammadali07/service-grap-go-api/services/acc/repository"
	"github.com/sirupsen/logrus"
)

type GLApp struct {
	repo *repository.GLRepository
	log  *logrus.Logger
}

func InitApp(repo *repository.GLRepository, log *logrus.Logger) *GLApp {
	return &GLApp{
		repo: repo,
		log:  log,
	}
}
