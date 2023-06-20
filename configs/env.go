package configs

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type env struct {
	DBHost     string
	DBUser     string
	DBName     string
	DBPassword string
}

func EnvConf() env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := env{
		DBHost:     os.Getenv("DBHost"),
		DBUser:     os.Getenv("DBUser"),
		DBName:     os.Getenv("DBName"),
		DBPassword: os.Getenv("DBPassword"),
	}

	return config

}
