package config

import "github.com/spf13/viper"

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	// Postgres configuration
	Host     string
	Port     string
	Username string
	Password string
	Database string
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
