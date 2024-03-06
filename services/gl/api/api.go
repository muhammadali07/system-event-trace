package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"ihsansolusi.co.id/information-centre/backend/app"
	"ihsansolusi.co.id/information-centre/backend/repository"
)

type GLApi struct {
	app       *app.ICApp
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
	setupEmployeeRoute(server, api)
	setupProjectRoute(server, api)
	setupActivityRoute(server, api)
}
