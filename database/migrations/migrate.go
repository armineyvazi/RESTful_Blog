package migrations

import (
	"log"

	"restful_blog/database"
	"restful_blog/models"
)

func MigrateModels() {
	database.InitializeDB()
	db := database.GetDB()
	err := db.AutoMigrate(&models.Category{}, &models.Post{})

	//Handle error
	if err != nil {
		log.Fatalf("migrate error%v", err)
	}
}
