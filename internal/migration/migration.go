package migration

import (
	"webgin/internal/pkg/models"
	"webgin/internal/services"
)

func Migrate() {
	services.DB.AutoMigrate(&models.User{})
}
