package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	// ==================================================== ROUTING ==================================================== //

	SERVERPort = GetEnv("SERVER_PORT", ":9070")
	AppPrefix  = GetEnv("BASE_URL", "/tkbai")
	JwtKey     = GetEnv("SV_JWT_KEY", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")
	DbUrl      = GetEnv("DB_URL", "tkbai:rytdin-ryqriN-kohqi6@tcp(127.0.0.1:3306)/tkbai?parseTime=true")
)

func GetEnv(key, fallback string) (value string) {
	logger := Log
	if value, ok := os.LookupEnv(key); ok {
		logger.Debug().Str(key, value).Msg("Env")
		return value
	}
	logger.Error().Str(key, fallback).Msg("Fallback")
	return fallback
}
