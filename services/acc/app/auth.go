package app

import (
	"net/http"

	"github.com/muhammadali07/system-event-trace/services/acc/pkg/utils"
)

func (a *AccountApp) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accountNo := r.Header.Get("account_no")
		pin := r.Header.Get("pin")
		// Lakukan verifikasi nomor rekening dan PIN di sini
		// Contoh: cek ke database
		encrypPin, err := utils.EncryptPin(pin)
		if err != nil {
			http.Error(w, "Failed to encrypt Pin", http.StatusUnauthorized)
			return
		}
		err = a.repo.VerifyAccount(accountNo, encrypPin)
		if err != nil {
			http.Error(w, "Autentikasi failed", http.StatusUnauthorized)
			return
		}

		// Lakukan verifikasi nomor rekening dan PIN di sini
		// Contoh: cek ke database

		next.ServeHTTP(w, r)
	})
}

func (a *AccountApp) AuthHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Lakukan autentikasi di sini menggunakan middleware
		a.AuthMiddleware(handler).ServeHTTP(w, r)
	}
}
