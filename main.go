package main

import (
	"giligili/conf"
	"giligili/server"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
