package configs

import (
	logger "github.com/Irajan/krane-stack/utils/logger"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type envConfigs struct {
	DbUser          string `env:"DB_USER_NAME"`
	DbPassword      string `env:"DB_PASSWORD"`
	DbHost          string `env:"DB_HOST"`
	DbName          string `env:"DB_NAME"`
	DbPort          string `env:"DB_PORT"`
	JwtAccessToken  string `env:"JWT_ACCESS_TOKEN_SECRET"`
	JwtRefreshToken string `env:"JWT_REFRESH_TOKEN_SECRET"`
}

var envConfig envConfigs

func initializeEnvConfigs() {

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading env file")
	}

	if err := env.Parse(&envConfig); err != nil {
		logger.Error("Error on parsing env file")
	}
}
