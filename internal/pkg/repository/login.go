package repository

import (
	"webgin/internal/pkg/models"
)

type Login interface {
	GetUser(username string) (*models.User, error)
}
