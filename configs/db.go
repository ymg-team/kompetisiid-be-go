// ref: https://onexlab-io.medium.com/golang-mysql-gorm-echo-feb526804c9f

package configs

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var envConfig Env = EnvConf()

var (
	DBUser     = envConfig.DBUser
	DBPassword = envConfig.DBPassword
	DBName     = envConfig.DBName
	DBHost     = envConfig.DBHost
	DBPort     = "3306"
	DBtype     = "mysql"
)

func GetDBType() string {
	return DBtype
}
