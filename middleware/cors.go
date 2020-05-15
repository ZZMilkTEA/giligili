package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = []string{"http://localhost:8493", "http://localhost:8080", "http://192.168.0.109:8493", "http://localhost:8081", "http://r3mix4lles.cn:9999",
		"http://r3mix4lles.cn", "http://182.92.212.142:9999"}
	config.AllowCredentials = true
	return cors.New(config)
}
