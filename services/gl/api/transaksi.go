package api

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"

	// "github.com/muhammadali07/service-grap-go-api/services/gl/api/serializer"
	"github.com/muhammadali07/service-grap-go-api/services/gl/models"
	utils "github.com/muhammadali07/service-grap-go-api/services/gl/pkg/utils"
)

var (
	DATE_FORMAT = "2006-01-02"
)

func (i *GLApi) createTransaction(ctx *fiber.Ctx) (err error) {
	var req models.Transaksi
	err = ctx.BodyParser(&req)
	if err != nil {
		remark := "failed to parse request body to JSON"
		i.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(remark)
		err = utils.HandleError(ctx, remark, http.StatusBadRequest)
		return
	}
	err = i.validator.Struct(req)
	if err != nil {
		remark := "failed to validate JSON request"
		i.log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error(remark)
		err = utils.HandleError(ctx, remark, http.StatusBadRequest)
		return
	}
	err = i.app.CreateTransaction(&req)
	if err != nil {
		err = utils.HandleError(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	err = utils.HandleSuccess(ctx, "create activity success", nil, http.StatusCreated)
	return
}

// func (i *ICAPI) getListActivities(ctx *fiber.Ctx) (err error) {
// 	var query serializer.QueryListActivities
// 	response := []serializer.ResponseListActivities{}
// 	err = ctx.QueryParser(&query)
// 	if err != nil {
// 		remark := "failed to parse query params to struct"
// 		i.log.WithFields(logrus.Fields{
// 			"error": err.Error(),
// 		}).Error(remark)
// 		err = handleError(ctx, remark, http.StatusBadRequest)
// 		return
// 	}
// 	err = i.validator.Struct(query)
// 	if err != nil {
// 		remark := "failed to validate query params"
// 		i.log.WithFields(logrus.Fields{
// 			"error": err.Error(),
// 		}).Error(remark)
// 		err = handleError(ctx, remark, http.StatusBadRequest)
// 		return
// 	}
// 	emps, err := i.app.GetListActivities(query.Start, query.End, query.Sort, query.ProjectDivsion, query.EmployeeID)
// 	if err != nil {
// 		err = handleError(ctx, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = copier.Copy(&response, emps)
// 	if err != nil {
// 		remark := "failed to parse query resposne"
// 		i.log.WithFields(logrus.Fields{
// 			"error": err.Error(),
// 		}).Error(remark)
// 		err = handleError(ctx, remark, http.StatusInternalServerError)
// 		return
// 	}
// 	data := groupResponseActivityByDate(response)
// 	err = handleSuccess(ctx, "query list employees success", data, http.StatusCreated)
// 	return
// }

// func groupResponseActivityByDate(activities []serializer.ResponseListActivities) (groupActivities map[string]interface{}) {
// 	groupActivities = make(map[string]interface{})
// 	today := time.Now().Format(DATE_FORMAT)
// 	keys := []string{}
// 	var list []serializer.ResponseListActivities
// 	for _, act := range activities {
// 		date := act.StartDate.Format(DATE_FORMAT)
// 		if date == today {
// 			date = "today"
// 		}
// 		_, ok := groupActivities[date]
// 		if !ok {
// 			list = []serializer.ResponseListActivities{}
// 			keys = append(keys, date)
// 		}
// 		list = append(list, act)
// 		groupActivities[date] = list
// 	}
// 	groupActivities["keys"] = keys
// 	return
// }

// func (i *ICAPI) updateActivity(ctx *fiber.Ctx) (err error) {
// 	var act models.Activity
// 	err = ctx.BodyParser(&act)
// 	if err != nil {
// 		remark := "failed to parse request body to JSON"
// 		i.log.WithFields(logrus.Fields{
// 			"error": err.Error(),
// 		}).Error(remark)
// 		err = handleError(ctx, remark, http.StatusBadRequest)
// 		return
// 	}
// 	err = i.validator.Struct(act)
// 	if err != nil {
// 		remark := "failed to validate JSON request"
// 		i.log.WithFields(logrus.Fields{
// 			"error": err.Error(),
// 		}).Error(remark)
// 		err = handleError(ctx, remark, http.StatusBadRequest)
// 		return
// 	}
// 	ID, err := ctx.ParamsInt("id")
// 	if err != nil {
// 		remark := "id param must be integer"
// 		i.log.WithFields(logrus.Fields{
// 			"error": err.Error(),
// 		}).Error(remark)
// 		err = handleError(ctx, remark, http.StatusBadRequest)
// 		return
// 	}
// 	err = i.app.UpdateActivity(uint(ID), act)
// 	if err != nil {
// 		err = handleError(ctx, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = handleSuccess(ctx, "update activity success", nil, http.StatusCreated)
// 	return
// }

// func (i *ICAPI) deleteActivity(ctx *fiber.Ctx) (err error) {
// 	ID, err := ctx.ParamsInt("id")
// 	if err != nil {
// 		remark := "id param must be integer"
// 		i.log.WithFields(logrus.Fields{
// 			"error": err.Error(),
// 		}).Error(remark)
// 		err = handleError(ctx, remark, http.StatusBadRequest)
// 		return
// 	}
// 	err = i.app.DeleteActivity(uint(ID))
// 	if err != nil {
// 		err = handleError(ctx, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = handleSuccess(ctx, "delete activity success", nil, http.StatusCreated)
// 	return
// }

func setupTransaksiRoute(server *fiber.App, api *GLApi) {
	group := server.Group("/journal")
	group.Post("", api.createTransaction)
	// group.Get("", api.getListActivities)
	// group.Put("/:id", api.updateActivity)
	// group.Delete("/:id", api.deleteActivity)
}
