package middleware

import (
	"net/http"
	"webgin/internal/config"
	"webgin/internal/constant"
	"webgin/internal/helpers"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func csrfMiddleware() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: config.Get(`server.secret.key`),
		ErrorFunc: func(c *gin.Context) {
			err := []string{"CSRF token mismatch"}
			helpers.FlashMessage(c, err, constant.ErrorKey)

			redirect := c.GetHeader("Referer")
			c.Redirect(http.StatusFound, redirect)
			c.Abort()
		},
	})
}
