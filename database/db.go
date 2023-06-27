package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeDB() {

	//Handle for test Bug
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	config.EnvConfigs.DB_USERNAME,
	//	config.EnvConfigs.DB_PASSWORD,
	//	config.EnvConfigs.DB_HOST,
	//	config.EnvConfigs.DB_PORT,
	//	config.EnvConfigs.DB_DATABASE,
	//)

	// TODO fix hardcode connection database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"A1r2m3i4n5@1234",
		"mysql",
		"3306",
		"go",
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Println("Connected to the database")

}

// GetDB returns the instance of the database connection
func GetDB() *gorm.DB {
	return db
}
