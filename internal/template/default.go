package template

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
)

// Default Template Middleware
func Default() *ginview.ViewEngine {
	tpl := ginview.New(goview.Config{
		Root:         "internal/views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        viewHelper(),
		DisableCache: true,
	})

	return tpl
}
