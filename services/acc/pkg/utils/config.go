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
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	Database         string
	Service          string
	KafkaHost        string
	KafkaPort        int
	KafkaTopic       string
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
		AppHost:          viper.GetString("APPHOST"),
		AppPort:          viper.GetString("APPPORT"),
		DatabasePort:     viper.GetInt("DATABASEPORT"),
		DatabaseHost:     viper.GetString("DATABASEHOST"),
		DatabaseUser:     viper.GetString("DATABASEUSER"),
		DatabasePassword: viper.GetString("DATABASEPASSWORD"),
		Database:         viper.GetString("DATABASE"),
		Service:          viper.GetString("SERVICE"),
		KafkaHost:        viper.GetString("KAFKAHOST"),
		KafkaPort:        viper.GetInt("KAFKAPORT"),
		KafkaTopic:       viper.GetString("KAFKATOPIC"),
	}

	return config, nil
}
