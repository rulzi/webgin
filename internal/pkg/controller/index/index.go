package controller

import (
	"net/http"

	"webgin/internal/constant"
	"webgin/internal/helpers"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Index interface {
	GetIndex(ctx *gin.Context)
	GetPageLogin(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type controller struct {
}

// IndexController init controller
func IndexController() Index {
	return &controller{}
}

func (c *controller) GetIndex(ctx *gin.Context) {
	//render with master
	ginview.HTML(ctx, http.StatusOK, "index", gin.H{
		"title": "Index title!",
	})
}

func (c *controller) GetPageLogin(ctx *gin.Context) {
	//render with master
	ginview.HTML(ctx, http.StatusOK, "index", gin.H{
		"title":   "User Login Page",
		"success": helpers.Flashes(ctx, constant.SuccessKey),
	})
}

func (c *controller) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete(constant.UserKey)
	session.Save()

	ctx.Redirect(http.StatusFound, helpers.BaseUrl()+"/")
}
