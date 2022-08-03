package controller

import (
	"net/http"
	"webgin/internal/constant"
	"webgin/internal/helpers"
	formvalidation "webgin/internal/pkg/form_validation"
	"webgin/internal/pkg/usecase"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type Register interface {
	GetRegister(ctx *gin.Context)
	PostRegister(ctx *gin.Context)
}

type controller struct {
	usecase usecase.Register
}

// RegisterController init controller
func RegisterController(usecase usecase.Register) Register {
	return &controller{
		usecase: usecase,
	}
}

func (c *controller) GetRegister(ctx *gin.Context) {
	//render with master
	ginview.HTML(ctx, http.StatusOK, "register", gin.H{
		"title":  "Register Page",
		"errors": helpers.Flashes(ctx, constant.ErrorKey),
		"csrf":   csrf.GetToken(ctx),
	})
}

func (c *controller) PostRegister(ctx *gin.Context) {

	valErr := helpers.ValidationErrors(ctx, &formvalidation.Register{})

	if valErr != nil {
		helpers.FlashMessage(ctx, valErr, constant.ErrorKey)

		redirect := ctx.GetHeader("Referer")
		ctx.Redirect(http.StatusFound, redirect)
		return
	}

	_, err := c.usecase.Register(ctx)

	if err != nil {
		helpers.FlashMessage(ctx, []string{err.Error()}, constant.ErrorKey)

		redirect := ctx.GetHeader("Referer")
		ctx.Redirect(http.StatusFound, redirect)
		return
	}

	helpers.FlashMessage(ctx, []string{"Registrasi Berhasil"}, constant.SuccessKey)
	ctx.Redirect(http.StatusFound, helpers.BaseUrl()+"/login")
}
