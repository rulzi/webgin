package middleware

import (
	"webgin/internal/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func session() gin.HandlerFunc {

	// store := cookie.NewStore([]byte(config.Get(`server.secret.key`)))
	store, _ := redis.NewStore(10, "tcp", config.Get(`redis.host`)+":"+config.Get(`redis.port`), "", []byte(config.Get(`server.secret.key`)))
	return sessions.Sessions("mysession", store)
}
