package config

import (
	"fmt"

	"github.com/usagifm/golang-jwt/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is creating a new connection to our databse
func SetupDatabaseConnection() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Failed to load ENV file")
	// }

	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbName := "golang-jwt"

	// trying to connect to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc-Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a Connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Diary{})
	return db

}

// disconnect to the database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")

	}

	dbSQL.Close()

}
