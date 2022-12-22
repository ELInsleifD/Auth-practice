package initializers

import "go-jwt-demo/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
