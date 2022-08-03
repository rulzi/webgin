package repository

import (
	"webgin/internal/pkg/models"
)

type Register interface {
	RegisterUser(user *models.User) (*models.User, error)
}
