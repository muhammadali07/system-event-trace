package http

import (
	"github.com/muhammadali07/service-grap-go-api/services/acc/internal/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine, uc usecase.AccountUseCase) {
	// Grup untuk API non-GraphQL
	api := r.Group("/api")

	// Registrasi nasabah
	api.POST("/register", handler.Register)

	// // Deposit
	// api.POST("/deposit", Handler.Deposit)

	// // Tarik tunai
	// api.POST("/withdraw", Handler.Withdraw)

	// // Transfer
	// api.POST("/transfer", Handler.Transfer)

	// // Saldo
	// api.GET("/saldo/:noRekening", Handler.GetSaldo)

	// // Mutasi
	// api.GET("/mutasi/:noRekening", Handler.GetMutasi)
}
