package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

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

func GenerateAccountNumber() string {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Generate 7 random digits.
	randomDigits := rand.Intn(10000000)
	accountNumber := fmt.Sprintf("320%07d", randomDigits)

	return accountNumber
}
