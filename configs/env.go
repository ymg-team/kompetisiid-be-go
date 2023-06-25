package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DBHost     string
	DBUser     string
	DBName     string
	DBPassword string
	MediaHost  string
}

func EnvConf() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Env{
		DBHost:     os.Getenv("DBHost"),
		DBUser:     os.Getenv("DBUser"),
		DBName:     os.Getenv("DBName"),
		DBPassword: os.Getenv("DBPassword"),
		MediaHost:  os.Getenv("MediaHost"),
	}

	return config

}
