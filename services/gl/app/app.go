package app

import (
	"github.com/sirupsen/logrus"
	"ihsansolusi.co.id/information-centre/backend/repository"
)

type GLApp struct {
	repo *repository.ICRepository
	log  *logrus.Logger
}

func InitApp(repo *repository.ICRepository, log *logrus.Logger) *ICApp {
	return &ICApp{
		repo: repo,
		log:  log,
	}
}
