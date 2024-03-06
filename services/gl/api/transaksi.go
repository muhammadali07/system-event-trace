package api

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	// "github.com/muhammadali07/service-grap-go-api/services/gl/api/serializer"
	"github.com/muhammadali07/service-grap-go-api/services/gl/models"
	utils "github.com/muhammadali07/service-grap-go-api/services/gl/pkg/utils"
)

var (
	DATE_FORMAT = "2006-01-02"
)

func (i *GLApi) createTransaction(ctx *fiber.Ctx) error {
	var req models.Transaksi
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

	err = i.app.CreateTransaction(&req)
	if err != nil {
		return utils.HandleError(ctx, err.Error(), http.StatusBadRequest)
	}

	return utils.HandleSuccess(ctx, "aktivitas berhasil dibuat", nil, http.StatusCreated)
}

func setupTransaksiRoute(server *fiber.App, api *GLApi) {
	group := server.Group("/journal")
	group.Post("", api.createTransaction)
	// Add other routes as needed
}
