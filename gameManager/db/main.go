package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//database global
var DB *gorm.DB

func SetupDB() *gorm.DB {

	//db config vars
	var dbHost string = os.Getenv("DB_HOST")
	var dbName string = os.Getenv("DB_NAME")
	var dbUser string = os.Getenv("DB_USERNAME")
	var dbPassword string = os.Getenv("DB_PASSWORD")
	var dbPort, err = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Println(psqlInfo)

	//connect to db
	db, dbError := gorm.Open("postgres", psqlInfo)
	if dbError != nil {
		panic("Failed to connect to database")
	}

	//fix for connection timeout
	//see: https://github.com/go-sql-driver/mysql/issues/257
	db.DB().SetMaxIdleConns(0)

	return db
}
