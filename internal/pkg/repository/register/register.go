package register

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
func Repository() repo.Register {
	return &repository{
		services.DB,
	}
}

func (repo *repository) RegisterUser(user *models.User) (*models.User, error) {
	err := repo.db.Save(user).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}
