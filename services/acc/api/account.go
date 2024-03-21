package api

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	utils "github.com/muhammadali07/service-grap-go-api/services/acc/pkg/utils"
)

var (
	DATE_FORMAT = "2006-01-02"
)

func (i *AcccountApi) createAccount(ctx *fiber.Ctx) error {
	var req models.Account
	err := ctx.BodyParser(&req)
	if err != nil {
		remark := "gagal mem-parsing body permintaan menjadi JSON"
		i.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(remark)
		return utils.HandleError(ctx, remark, http.StatusBadRequest)
	}

	err = i.validator.Struct(req)
	if err != nil {
		remark := "gagal memvalidasi permintaan JSON"
		i.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(remark)
		return utils.HandleError(ctx, remark, http.StatusBadRequest)
	}

	res, err := i.app.CreateAccount(&req)
	if err != nil {
		return utils.HandleError(ctx, err.Error(), http.StatusBadRequest)
	}

	out_response := map[string]interface{}{
		"nomor_rekening_nasabah": res,
		"data":                   req,
	}

	return utils.HandleSuccess(ctx, "registrasi akun berhasil dibuat", out_response, http.StatusCreated)
}

func setupAccountRoute(server *fiber.App, api *AcccountApi) {
	group := server.Group("/account")
	group.Post("/daftar", api.createAccount)
	// Add other routes as needed
}
