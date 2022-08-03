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

type Login interface {
	GetLogin(ctx *gin.Context)
	PostLogin(ctx *gin.Context)
}

type controller struct {
	usecase usecase.Login
}

// LoginController init controller
func LoginController(usecase usecase.Login) Login {
	return &controller{
		usecase: usecase,
	}
}

func (c *controller) GetLogin(ctx *gin.Context) {
	//render with master
	ginview.HTML(ctx, http.StatusOK, "login", gin.H{
		"title":   "Login Page",
		"errors":  helpers.Flashes(ctx, constant.ErrorKey),
		"success": helpers.Flashes(ctx, constant.SuccessKey),
		"csrf":    csrf.GetToken(ctx),
	})
}

func (c *controller) PostLogin(ctx *gin.Context) {
	valErr := helpers.ValidationErrors(ctx, &formvalidation.Login{})

	if valErr != nil {
		helpers.FlashMessage(ctx, valErr, constant.ErrorKey)

		redirect := ctx.GetHeader("Referer")
		ctx.Redirect(http.StatusFound, redirect)
		return
	}

	_, err := c.usecase.Login(ctx)

	if err != nil {
		helpers.FlashMessage(ctx, []string{err.Error()}, constant.ErrorKey)

		redirect := ctx.GetHeader("Referer")
		ctx.Redirect(http.StatusFound, redirect)
		return
	}

	helpers.FlashMessage(ctx, []string{"Login Berhasil"}, constant.SuccessKey)
	ctx.Redirect(http.StatusFound, helpers.BaseUrl()+"/")
}
