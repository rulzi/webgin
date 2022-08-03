package template

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

// Custom Template Middleware
func Web() gin.HandlerFunc {
	tpl := ginview.NewMiddleware(goview.Config{
		Root:         "internal/views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        viewHelper(),
		DisableCache: true,
	})

	return tpl
}
