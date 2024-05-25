package utils

import (
	"github.com/spf13/viper"
)

// Config adalah struct yang akan digunakan untuk menyimpan konfigurasi aplikasi.
type Config struct {
	// Tambahkan field konfigurasi sesuai dengan kebutuhan Anda.
	// Misalnya:
	DefaultPort              string
	KafkaHost                string
	KafkaPort                int
	KafkaServiceName         string
	DatabaseHost             string
	DatabasePort             int
	DatabaseUser             string
	DatabasePassword         string
	DatabaseDriver           string
	DatabaseSchema           map[string]string
	Database                 string
	TelemetryEndpoint        string
	DefaultTelemetryEndpoint string
}

// InitConfig inisialisasi konfigurasi menggunakan Viper.
func InitConfig() (*Config, error) {

	// // Set path file konfigurasi untuk Viper
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Baca konfigurasi dari file (opsional)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	DB_SCHEMAS := map[string]string{
		"core": viper.GetString("CORE_SCHEMA"),
	}

	// Buat instance Config dan isi sesuai dengan konfigurasi yang dibaca
	config := &Config{
		DefaultPort:              viper.GetString("DEFAULT_PORT"),
		KafkaHost:                viper.GetString("KAFKA_HOST"),
		KafkaPort:                viper.GetInt("KAFKA_PORT"),
		KafkaServiceName:         viper.GetString("KAFKA_SERVICE"),
		DatabasePort:             viper.GetInt("DATABASE_PORT"),
		DatabaseHost:             viper.GetString("DATABASE_HOST"),
		DatabaseUser:             viper.GetString("DATABASE_USER"),
		DatabasePassword:         viper.GetString("DATABASE_PASSWORD"),
		DatabaseDriver:           viper.GetString("DATABASE_DRIVER"),
		DatabaseSchema:           DB_SCHEMAS,
		Database:                 viper.GetString("DATABASE"),
		TelemetryEndpoint:        viper.GetString("TELEMETRY_ENDPOINT"),
		DefaultTelemetryEndpoint: viper.GetString("DEFAULT_TELEMETRY_ENDPOINT"),
	}

	return config, nil
}
