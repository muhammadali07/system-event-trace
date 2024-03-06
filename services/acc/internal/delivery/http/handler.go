package http

import (
	"app/internal/domain"
	"app/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	uc usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Remark: err.Error(),
		})
		return
	}

	// Validasi input
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Remark: err.Error(),
		})
		return
	}

	// Hash PIN
	hashedPIN, err := bcrypt.GenerateFromPassword([]byte(req.PIN), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{
			Remark: "Gagal generate hash PIN",
		})
		return
	}

	// Buat akun baru
	account, err := h.uc.CreateAccount(c.Request.Context(), req.Nama, req.NIK, req.NoHP, hashedPIN)
	if err != nil {
		switch err {
		case domain.ErrDuplicateNIK:
			c.JSON(http.StatusBadRequest, errorResponse{
				Remark: "NIK sudah terdaftar",
			})
		case domain.ErrDuplicateNoHP:
			c.JSON(http.StatusBadRequest, errorResponse{
				Remark: "Nomor HP sudah terdaftar",
			})
		default:
			c.JSON(http.StatusInternalServerError, errorResponse{
				Remark: "Gagal create akun",
			})
		}
		return
	}

	// Kirim response
	c.JSON(http.StatusCreated, registerResponse{
		NoRekening: account.NoRekening,
	})
}

func (h *Handler) Deposit(c *gin.Context) {
	var req depositRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Remark: err.Error(),
		})
		return
	}

	// Validasi input
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{
			Remark: err.Error(),
		})
		return
	}

	// Lakukan deposit
	saldo, err := h.uc.Deposit(c.Request.Context(), req.NoRekening, req.Nominal)
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			c.JSON(http.StatusBadRequest, errorResponse{
				Remark: "Nomor rekening tidak ditemukan",
			})
		default:
			c.JSON(http.StatusInternalServerError, errorResponse{
				Remark: "Gagal deposit",
			})
		}
		return
	}

	// Kirim response
	c.JSON(http.StatusOK, depositResponse{
		Saldo: saldo,
	})
}

// func (h *Handler) Withdraw(c *gin.Context) {
//   var req withdrawRequest
//   if err := c.ShouldBindJSON(&req); err != nil {
//     c.JSON(http.StatusBadRequest, errorResponse{
//       Remark: err.Error(),
//     })
//     return
//   }

//   // Validasi input
//   if err := validate.Struct(req); err != nil {
//     c.JSON(http.StatusBadRequest, errorResponse{
//       Remark: err.Error(),
//     })
//     return
//   }

//   // Lakukan tarik tunai
//   saldo, err := h.uc.Withdraw(c.Request.Context(), req.NoRekening, req.Nominal)
//   if err != nil {
//     switch err {
//     case domain.ErrAccountNotFound:
//       c.JSON(http.StatusBadRequest, errorResponse{
//         Remark: "Nomor rekening tidak ditemukan",
//       })
//     case domain.ErrInsufficientSaldo:
//       c.JSON(http.StatusBadRequest
