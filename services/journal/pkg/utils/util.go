package utils

import (
	"sync"

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

func GetNextNumber() int {
	counter := 0
	var mutex sync.Mutex // Mutex for thread-safe access

	mutex.Lock()         // Acquire the mutex lock before accessing the counter
	defer mutex.Unlock() // Release the mutex lock after accessing the counter

	counter++ // Increment the counter
	return counter
}
