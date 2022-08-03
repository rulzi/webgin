package helpers

import (
	"errors"
	"log"
	"os"
	"webgin/internal/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func BaseUrl() string {
	return config.Get(`appurl`)
}

func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return string(result)
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

func DumpAndDie(data ...interface{}) {
	spew.Dump(data)

	os.Exit(1)
}

func FlashMessage(c *gin.Context, messages []string, vars ...string) {
	session := sessions.Default(c)
	for _, message := range messages {
		session.AddFlash(message, vars...)
	}
	if err := session.Save(); err != nil {
		log.Printf("error in flashMessage saving session: %s", err)
	}
}

func Flashes(c *gin.Context, vars ...string) []interface{} {
	session := sessions.Default(c)
	flashes := session.Flashes(vars...)
	if len(flashes) != 0 {
		if err := session.Save(); err != nil {
			log.Printf("error in flashes saving session: %s", err)
		}
	}
	return flashes
}

func GetErrorMsg(field string, fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " should be email format"
	case "lte":
		return field + " should be less than " + fe.Param()
	case "gte":
		return field + " should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func ValidationErrors(c *gin.Context, m interface{}) []string {
	err := c.ShouldBind(m)

	if err != nil {
		var ve validator.ValidationErrors
		var errMsg []string
		if errors.As(err, &ve) {

			for _, fe := range ve {
				errMsg = append(errMsg, GetErrorMsg(fe.Field(), fe))
			}

			return errMsg
		}
	}

	return nil
}
