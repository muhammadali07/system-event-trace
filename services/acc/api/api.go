package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadali07/system-event-trace/services/acc/app"
	"github.com/muhammadali07/system-event-trace/services/acc/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AcccountApi struct {
	app       *app.AccountApp
	log       *logrus.Logger
	validator *validator.Validate
}

func InitServer(server *fiber.App, db *gorm.DB, log *logrus.Logger, validator *validator.Validate) {
	repo := repository.InitRepository(db, log)
	app := app.InitApp(repo, log)
	api := &AcccountApi{
		app:       app,
		log:       log,
		validator: validator,
	}
	setupAccountRoute(server, api)
	setupTransaksiRoute(server, api)
}
