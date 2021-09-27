package util

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/lib/pq"
	"github.com/spf13/viper"
)

// Configurations wraps all the config variables required by the auth service
type Configurations struct {
	// ServerAddress string
	ServerPort string
	DBHost     string
	DBName     string
	DBUser     string
	DBPass     string
	DBPort     string
	DBConn     string
}

// NewConfigurations returns a new Configuration object
func NewConfigurations(logger log.Logger) *Configurations {

	viper.AutomaticEnv()

	dbURL := viper.GetString("DATABASE_URL")
	conn, _ := pq.ParseURL(dbURL)
	level.Debug(logger).Log("msg", "found database url in env, connection string is formed by parsing it")
	level.Debug(logger).Log("db connection string", conn)

	// viper.SetDefault("SERVER_ADDRESS", "0.0.0.0:8088")
	viper.SetDefault("SERVER_PORT", "8089")
	viper.SetDefault("DB_HOST", "sadhelx_db")
	// viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "sdx_sertifikasi_db")
	viper.SetDefault("DB_USER", "sadhelx_usr")
	viper.SetDefault("DB_PASSWORD", "s4dhelx")
	viper.SetDefault("DB_PORT", "5432")
	// viper.SetDefault("DB_PORT", "54321")

	configs := &Configurations{
		// ServerAddress: viper.GetString("SERVER_ADDRESS"),
		ServerPort: viper.GetString("SERVER_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPass:     viper.GetString("DB_PASSWORD"),
		DBPort:     viper.GetString("DB_PORT"),
		DBConn:     conn,
	}

	level.Debug(logger).Log("serve port", configs.ServerPort)
	level.Debug(logger).Log("db host", configs.DBHost)
	level.Debug(logger).Log("db name", configs.DBName)
	level.Debug(logger).Log("db port", configs.DBPort)

	return configs
}
