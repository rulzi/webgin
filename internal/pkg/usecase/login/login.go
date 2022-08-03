package login

import (
	"webgin/internal/constant"
	"webgin/internal/helpers"
	"webgin/internal/pkg/models"
	"webgin/internal/pkg/repository"
	uc "webgin/internal/pkg/usecase"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type usecase struct {
	repo repository.Login
}

// Usecase init usecase
func Usecase(repository repository.Login) uc.Login {
	return &usecase{
		repo: repository,
	}
}

func (uc *usecase) Login(c *gin.Context) (*models.User, error) {
	username := c.PostForm("username")

	user, err := uc.repo.GetUser(username)

	if err != nil {
		return nil, err
	}

	if err := helpers.ComparePassword(user.Password, c.PostForm("password")); err != nil {
		return nil, err
	}

	session := sessions.Default(c)
	session.Set(constant.UserKey, user.Username)

	if err := session.Save(); err != nil {
		return nil, err
	}

	return user, nil
}
