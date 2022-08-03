package register

import (
	"webgin/internal/helpers"
	"webgin/internal/pkg/models"
	"webgin/internal/pkg/repository"
	uc "webgin/internal/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type usecase struct {
	repo repository.Register
}

// Usecase init usecase
func Usecase(repository repository.Register) uc.Register {
	return &usecase{
		repo: repository,
	}
}

func (uc *usecase) Register(c *gin.Context) (*models.User, error) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	username := c.PostForm("username")
	address := c.PostForm("address")
	password := helpers.HashPassword(c.PostForm("password"))

	user := models.User{
		Email:    email,
		Name:     name,
		Username: username,
		Address:  address,
		Password: password,
	}

	userSave, err := uc.repo.RegisterUser(&user)

	if err != nil {
		return nil, err
	}

	return userSave, nil
}
