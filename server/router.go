package server

import (
	"giligili/api"
	"giligili/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.LoggerToFile())
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user", api.UserRegister)

		// 用户登录
		v1.POST("login", api.UserLogin)

		v1.GET("verify", api.Verify)
		v1.GET("refresh", api.Refresh)
		v1.GET("sayHello", api.SayHello)
		// 用户登出
		v1.GET("logout", api.UserLogout)

		//获取用户信息
		v1.GET("user/:id", api.GetUser)
		v1.GET("users", api.ListUser)
		// 需要登录保护的
		authedUser := v1.Group("/")
		authedUser.Use(middleware.AuthUserRequired())
		{
			// User Routing
			authedUser.DELETE("user/:id/logout", api.UserLogout)
			// 视频操作
			authedUser.POST("videos", api.CreateVideo)
			authedUser.PUT("video/:id", api.UpdateVideo)
			authedUser.DELETE("video/:id", api.DeleteVideo)
		}
		// 需要验证审查员身份的
		authAdmin := v1.Group("/")
		authAdmin.Use(middleware.AuthInspectorRequired())
		{
			v1.PUT("user/change-permission", api.ChangeUserPermission)
			v1.DELETE("user/:id", api.DeleteUser)
		}

		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		v1.GET("user/:id/videos", api.ListVideoByUser)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 其他
		v1.POST("upload/token", api.UploadToken)
	}

	// swagger文档
	// 游览器打开 http://localhost:3000/swagger/index.html
	r.StaticFile("/swagger.json", "./swagger/swagger.json")
	r.Static("/swagger", "./swagger/dist")

	return r
}
