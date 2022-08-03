package usecase

import (
	"webgin/internal/pkg/models"

	"github.com/gin-gonic/gin"
)

type Login interface {
	Login(c *gin.Context) (*models.User, error)
}
