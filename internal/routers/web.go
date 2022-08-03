package routers

import (
	"context"
	"webgin/internal/middleware"
	"webgin/internal/routers/index"
	"webgin/internal/routers/login"
	"webgin/internal/routers/register"

	"github.com/gin-gonic/gin"
)

func Web(ctx context.Context, r *gin.Engine) *gin.Engine {

	indexController := index.Controller()
	loginController := login.Controller()
	registerController := register.Controller()

	g := r.Group("/")
	g.Use(middleware.GuestMiddleware)
	{
		g.GET("/login", loginController.GetLogin)
		g.POST("/login", loginController.PostLogin)

		g.GET("/register", registerController.GetRegister)
		g.POST("/register", registerController.PostRegister)
	}

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware)
	{
		auth.GET("/", indexController.GetPageLogin)
		auth.GET("/logout", indexController.Logout)
	}

	return r
}
