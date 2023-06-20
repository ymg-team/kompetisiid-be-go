// ref: https://onexlab-io.medium.com/golang-mysql-gorm-echo-feb526804c9f

package configs

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var envConfig env = EnvConf()

var (
	DBUser     = envConfig.DBUser
	DBPassword = envConfig.DBPassword
	DBName     = envConfig.DBName
	DBHost     = envConfig.DBHost
	DBPort     = "3306"
	DBtype     = "mysql"
)

var DB *gorm.DB

func GetDBType() string {
	return DBtype
}

func GetMySQLConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return dataBase
}

func ConnectDB() *gorm.DB {
	var err error
	conString := GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}
