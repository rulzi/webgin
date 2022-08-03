package providers

import (
	"context"
	"net/http"
	"webgin/internal/config"
	"webgin/internal/middleware"
	"webgin/internal/routers"
	"webgin/internal/template"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func Route(ctx context.Context) *gin.Engine {
	if !config.GetBoolean(`debug`) {
		gin.SetMode(gin.ReleaseMode)
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Static File
	r.Static("/assets", "./assets")

	//new default template Engine
	r.HTMLRender = template.Default()

	// Notfound
	r.NoRoute(func(c *gin.Context) {
		ginview.HTML(c, http.StatusNotFound, "errors/404.html", gin.H{})
	})

	// Middleware
	r = middleware.Middleware(ctx, r)

	// Routes Web
	r = routers.Web(ctx, r)

	return r
}
