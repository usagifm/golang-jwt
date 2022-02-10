package config

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSetupDatabaseConnection(t *testing.T) {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	t.Fatal("Failed to load ENV file")
	// }

	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbName := "golang-jwt"

	// trying to connect to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc-Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("Failed to create a Connection to database")
	}

	fmt.Print(db)

}
