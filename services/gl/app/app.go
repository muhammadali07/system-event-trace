package app

import (
	"github.com/sirupsen/logrus"
	"github.com/muhammadali07/service-grap-go-api/services/gl/repository"
)

type GLApp struct {
	repo *repository.ICRepository
	log  *logrus.Logger
}

func InitApp(repo *repository.ICRepository, log *logrus.Logger) *GLApp {
	return &GLApp{
		repo: repo,
		log:  log,
	}
}
