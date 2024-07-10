package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVICE        string
	USER_ROUTER         string
	RESERVATION_SERVICE string
	PAYMENT_SERVICE     string
	API_GATEWAY         string
	DB_HOST             string
	DB_PORT             string
	DB_USER             string
	DB_PASSWORD         string
	DB_NAME             string
	SIGNING_KEY         string
}

func Load() Config {
	if err := godotenv.Load("/home/azizbek/github/Restaurant/SitAndEat_Payment_service/.env"); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}
	config.USER_SERVICE = cast.ToString(Coalesce("USER_SERVICE", "50051"))
	config.USER_ROUTER = cast.ToString(Coalesce("USER_ROUTER", "50052"))
	config.RESERVATION_SERVICE = cast.ToString(Coalesce("RESERVATION_SERVICE", "50053"))
	config.PAYMENT_SERVICE = cast.ToString(Coalesce("PAYMENT_SERVICE", "50054"))
	config.API_GATEWAY = cast.ToString(Coalesce("API_GATEWAY", "50055"))
	config.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(Coalesce("DB_PORT", "5432"))
	config.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(Coalesce("DB_PASSWORD", "root"))
	config.DB_NAME = cast.ToString(Coalesce("DB_NAME", "paymentservice"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", "secret"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
