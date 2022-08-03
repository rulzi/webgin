package template

import (
	"html/template"
	"webgin/internal/helpers"
)

func viewHelper() template.FuncMap {
	fc := template.FuncMap{}

	fc["baseurl"] = func() string {
		return helpers.BaseUrl()
	}

	return fc
}
