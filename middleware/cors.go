package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Kind", "Cookie"}
	str := os.Getenv("CORS")
	strs := strings.Split(str, ",")
	config.AllowOrigins = strs
	config.AllowCredentials = true
	return cors.New(config)
}
