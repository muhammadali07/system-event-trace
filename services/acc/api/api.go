package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadali07/service-grap-go-api/services/acc/app"
	"github.com/muhammadali07/service-grap-go-api/services/acc/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GLApi struct {
	app       *app.GLApp
	log       *logrus.Logger
	validator *validator.Validate
}

func InitServer(server *fiber.App, db *gorm.DB, log *logrus.Logger, validator *validator.Validate) {
	repo := repository.InitRepository(db, log)
	app := app.InitApp(repo, log)
	api := &GLApi{
		app:       app,
		log:       log,
		validator: validator,
	}
	setupTransaksiRoute(server, api)
}
