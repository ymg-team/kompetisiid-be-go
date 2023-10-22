package configs

import (
	"os"
)

// err := godotenv.Load()
// if err != nil {
// 	log.Fatal("Error loading .env file")
// }
var TELEGRAM_KI_TOKEN = os.Getenv("TELEGRAM_KI_TOKEN")
var TELEGRAM_KI_CHANNEL = os.Getenv("TELEGRAM_KI_CHANNEL")

// func EnvTelegram() (string, string) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	return os.Getenv("MONGO_URI"), os.Getenv("")
// }

// func EnvMongoDB() string {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	return os.Getenv("MONGO_DB")
// }
