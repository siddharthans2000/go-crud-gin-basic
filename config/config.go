package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func getEnvVariable(path string) string {
	value := os.Getenv(path)
	if value != "" {
		log.Fatal("The env variable does not exist")
	}
	return value
}

func Connect() {
	// mysqldb mentioned here is the name of the container we use for mysql db
	d, err := gorm.Open(mysql.Open(getEnvVariable("MYSQL_ROOT_USERNAME")+":"+getEnvVariable("MYSQL_ROOT_PASSWORD")+
		"@tcp(mysqldb)/"+getEnvVariable("MYSQL_DATABASE")+"?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		log.Fatal("Issue with connecting the DB")
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
