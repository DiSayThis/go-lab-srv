package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Http HttpConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type HttpConfig struct {
	Port string
	Host string
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error file .env", err.Error())
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Http: HttpConfig{
			Port: port,
			Host: os.Getenv("HOST"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
