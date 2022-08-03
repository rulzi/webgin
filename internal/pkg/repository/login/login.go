package login

import (
	"webgin/internal/pkg/models"
	repo "webgin/internal/pkg/repository"
	"webgin/internal/services"

	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

// Repository init usecase
func Repository() repo.Login {
	return &repository{
		services.DB,
	}
}

func (repo *repository) GetUser(username string) (*models.User, error) {
	var user models.User

	err := repo.db.Where("username = ?", username).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
