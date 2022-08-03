package middleware

import (
	"net/http"
	"webgin/internal/constant"
	"webgin/internal/helpers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(constant.UserKey)
	if user == nil {

		c.Redirect(http.StatusFound, helpers.BaseUrl()+"/login")
		c.Abort()
		return
	}
	c.Next()
}

func GuestMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(constant.UserKey)
	if user != nil {

		c.Redirect(http.StatusFound, helpers.BaseUrl()+"/")
		c.Abort()
		return
	}
	c.Next()
}
