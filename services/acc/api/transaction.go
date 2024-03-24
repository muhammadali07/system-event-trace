package api

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/muhammadali07/service-grap-go-api/services/acc/models"
	utils "github.com/muhammadali07/service-grap-go-api/services/acc/pkg/utils"
)

func (i *AcccountApi) cashDeposit(ctx *fiber.Ctx) error {
	var req models.TransactionDepositWithdraw
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

	res, err := i.app.CashDeposit(req)
	if err != nil {
		return utils.HandleError(ctx, err.Error(), http.StatusBadRequest)
	}

	out_response := map[string]interface{}{
		"balance": res,
	}

	return utils.HandleSuccess(ctx, "cash deposito success", out_response, http.StatusCreated)
}

func (i *AcccountApi) cashWithdraw(ctx *fiber.Ctx) error {
	var req models.TransactionDepositWithdraw
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

	res, err := i.app.CashWithDraw(req)
	if err != nil {
		return utils.HandleError(ctx, err.Error(), http.StatusBadRequest)
	}

	out_response := map[string]interface{}{
		"balance": res,
	}

	return utils.HandleSuccess(ctx, "cash withdraw success", out_response, http.StatusCreated)
}

func (i *AcccountApi) transferKliring(ctx *fiber.Ctx) error {
	var req models.TransactionKliring
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

	res, err := i.app.TransferKliring(req)
	if err != nil {
		return utils.HandleError(ctx, err.Error(), http.StatusBadRequest)
	}

	out_response := map[string]interface{}{
		"balance": res,
	}

	return utils.HandleSuccess(ctx, "cash withdraw success", out_response, http.StatusCreated)
}

func (i *AcccountApi) getAccountBalance(ctx *fiber.Ctx) error {
	accountNumber := ctx.Params("accountNumber")

	i.log.WithFields(logrus.Fields{"accountNumber": accountNumber})

	res, err := i.app.GetAccountBalance(accountNumber)
	if err != nil {
		return utils.HandleError(ctx, err.Error(), http.StatusBadRequest)
	}

	out_response := map[string]interface{}{
		"saldo": res.Respdata,
	}

	return utils.HandleSuccess(ctx, res.RespMsg, out_response, http.StatusCreated)
}

func setupTransaksiRoute(server *fiber.App, api *AcccountApi) {
	group := server.Group("/transaction")
	group.Post("/tabung", api.cashDeposit)
	group.Post("/tarik", api.cashWithdraw)
	group.Post("/transfer", api.transferKliring)
	group.Get("/cek-saldo/:accountNumber", api.getAccountBalance)
	// Add other routes as needed
}
