package db

import (
	"fmt"
	"ki-be/configs"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetMySQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.DBUser,
		configs.DBPassword,
		configs.DBHost,
		configs.DBPort,
		configs.DBName)

	return dataBase
}

func ConnectDB() *gorm.DB {
	var err error
	conString := GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(configs.GetDBType(), conString)
	DB.SingularTable(true)
	if err != nil {
		log.Panic(err)
	}

	return DB
}
