package utils

import "github.com/gofiber/fiber/v2"

func HandleError(ctx *fiber.Ctx, remark string, status int) error {
	response := make(map[string]any)
	response["remark"] = remark
	ctx.Status(status)
	return ctx.JSON(response)
}

func HandleSuccess(ctx *fiber.Ctx, remark string, data any, status int) error {
	response := make(map[string]any)
	response["remark"] = remark
	response["data"] = data
	ctx.Status(status)
	return ctx.JSON(response)
}
