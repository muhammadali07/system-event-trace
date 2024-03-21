package utils

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

// encryptPin digunakan untuk mengenkripsi pin sebelum disimpan ke dalam database
func EncryptPin(pin string) (string, error) {
	// Generate salt
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Hash pin dengan bcrypt dan salt
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Gabungkan salt dan hash menjadi string yang dihex
	encryptedPin := hex.EncodeToString(append(salt, hashedPin...))

	return encryptedPin, nil
}

// verifyPin memverifikasi apakah pin yang dimasukkan oleh pengguna cocok dengan pin yang dienkripsi
func VerifyPin(plainPin, encryptedPin string) bool {
	// Decode hasil enkripsi menjadi byte array
	decodedHash, err := hex.DecodeString(encryptedPin)
	if err != nil {
		return false
	}

	// Membandingkan kata sandi mentah dengan hasil enkripsi
	err = bcrypt.CompareHashAndPassword(decodedHash, []byte(plainPin))
	return err == nil
}
