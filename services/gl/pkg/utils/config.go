package utils

import (
	"github.com/spf13/viper"
)

// Config adalah struct yang akan digunakan untuk menyimpan konfigurasi aplikasi.
type Config struct {
	// Tambahkan field konfigurasi sesuai dengan kebutuhan Anda.
	// Misalnya:
	AppHost          string
	AppPort          string
	Driver           string
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	Database         string
	Service          string
}

// InitConfig inisialisasi konfigurasi menggunakan Viper.
func InitConfig() (*Config, error) {
	// Inisialisasi Viper
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Baca konfigurasi dari file (opsional)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Buat instance Config dan isi sesuai dengan konfigurasi yang dibaca
	config := &Config{
		AppHost:          viper.GetString("AppHost"),
		AppPort:          viper.GetString("AppPort"),
		Driver:           viper.GetString("Driver"),
		DatabasePort:     viper.GetInt("DatabasePort"),
		DatabaseHost:     viper.GetString("DatabaseHost"),
		DatabaseUser:     viper.GetString("DatabaseUser"),
		DatabasePassword: viper.GetString("DatabasePassword"),
		Database:         viper.GetString("Database"),
		Service:          viper.GetString("Service"),
	}

	return config, nil
}
