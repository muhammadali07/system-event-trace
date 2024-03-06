package http

import (
	"net/http"

	"github.com/muhammadali07/service-grap-go-api/services/acc/internal/domain"
	"github.com/muhammadali07/service-grap-go-api/services/acc/internal/usecase"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	uc usecase.AccountUseCase
}

func NewHandler(uc usecase.AccountUseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Register(c *gin.Context) {
	var req domain.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Remark: err.Error(),
		})
		return
	}

	// Validasi input
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Remark: err.Error(),
		})
		return
	}

	// Hash PIN
	hashedPIN, err := bcrypt.GenerateFromPassword([]byte(req.PIN), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Remark: "Gagal generate hash PIN",
		})
		return
	}

	// Buat akun baru
	paramsRegist := domain.RegisterRequest{
		Nama: req.Nama,
		NIK:  req.NIK,
		NoHP: req.NoHP,
		PIN:  string(hashedPIN),
	}
	account, err := h.uc.CreateAccount(c.Request.Context(), paramsRegist)
	if err != nil {
		switch err {
		case domain.ErrDuplicateNIK:
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Remark: "NIK sudah terdaftar",
			})
		case domain.ErrDuplicateNoHP:
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Remark: "Nomor HP sudah terdaftar",
			})
		default:
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Remark: "Gagal create akun",
			})
		}
		return
	}

	// Kirim response
	c.JSON(http.StatusCreated, domain.RegisterResponse{
		NoRekening: account.NoRekening,
	})
}

// func (h *Handler) Deposit(c *gin.Context) {
// 	var req domain.DepositRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
// 			Remark: err.Error(),
// 		})
// 		return
// 	}

// 	// Validasi input
// 	if err := validate.Struct(req); err != nil {
// 		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
// 			Remark: err.Error(),
// 		})
// 		return
// 	}

// 	// Lakukan deposit
// 	saldo, err := h.uc.Deposit(c.Request.Context(), req.NoRekening, req.Nominal)
// 	if err != nil {
// 		switch err {
// 		case domain.ErrAccountNotFound:
// 			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
// 				Remark: "Nomor rekening tidak ditemukan",
// 			})
// 		default:
// 			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
// 				Remark: "Gagal deposit",
// 			})
// 		}
// 		return
// 	}

// 	// Kirim response
// 	c.JSON(http.StatusOK, domain.DepositResponse{
// 		Saldo: saldo,
// 	})
// }

// func (h *Handler) Withdraw(c *gin.Context) {
//   var req withdrawRequest
//   if err := c.ShouldBindJSON(&req); err != nil {
//     c.JSON(http.StatusBadRequest, domain.ErrorResponse{
//       Remark: err.Error(),
//     })
//     return
//   }

//   // Validasi input
//   if err := validate.Struct(req); err != nil {
//     c.JSON(http.StatusBadRequest, domain.ErrorResponse{
//       Remark: err.Error(),
//     })
//     return
//   }

//   // Lakukan tarik tunai
//   saldo, err := h.uc.Withdraw(c.Request.Context(), req.NoRekening, req.Nominal)
//   if err != nil {
//     switch err {
//     case domain.ErrAccountNotFound:
//       c.JSON(http.StatusBadRequest, domain.ErrorResponse{
//         Remark: "Nomor rekening tidak ditemukan",
//       })
//     case domain.ErrInsufficientSaldo:
//       c.JSON(http.StatusBadRequest
