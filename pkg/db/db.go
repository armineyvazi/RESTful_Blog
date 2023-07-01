package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func InitializeDB(config map[string]string) {

	// Create dsn to connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["USER"],
		config["PASSWORD"],
		config["HOST"],
		config["PORT"],
		config["DBNAME"],
	)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	log.Println("Connected to the database")
}

// GetDB returns the instance of the database connection
func GetDB() *gorm.DB {
	return Db
}

func CloseDB() error {
	// Close the database connection
	db, err := Db.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
