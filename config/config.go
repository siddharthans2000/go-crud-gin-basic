package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func getEnvVariable(path string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error getting env file")
	}
	value := os.Getenv(path)
	return value
}

func Connect() {
	// mysqldb mentioned here is the name of the container we use for mysql db
	dsn := getEnvVariable("MYSQL_ROOT_USERNAME") + ":" + getEnvVariable("MYSQL_ROOT_PASSWORD") +
		"@tcp(172.17.0.2:3306)/" + getEnvVariable("MYSQL_DATABASE") + "?charset=utf8&parseTime=True"
	// log.Fatalln(dsn)
	fmt.Println(dsn)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Issue with connecting the DB")

	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
