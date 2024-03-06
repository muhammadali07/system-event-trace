package http

import (
	"app/internal/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, uc usecase.UseCase) {
	// Grup untuk API non-GraphQL
	api := r.Group("/api")

	// Registrasi nasabah
	api.POST("/register", Handler.Register)

	// Deposit
	api.POST("/deposit", Handler.Deposit)

	// Tarik tunai
	api.POST("/withdraw", Handler.Withdraw)

	// Transfer
	api.POST("/transfer", Handler.Transfer)

	// Saldo
	api.GET("/saldo/:noRekening", Handler.GetSaldo)

	// Mutasi
	api.GET("/mutasi/:noRekening", Handler.GetMutasi)
}
