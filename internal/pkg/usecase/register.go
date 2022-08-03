package usecase

import (
	"webgin/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

type Register interface {
	Register(c *gin.Context) (*models.User, error)
}
