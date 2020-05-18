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
		//---------------------游客可使用的接口-----------------------
		v1.GET("videos/:id", api.ShowVideo)
		v1.GET("videos/:id/comments", api.ListCommentsByMediaId)

		v1.GET("passed-videos", api.ListPassedVideos)
		v1.GET("user/:id/passed-videos", api.ListPassedVideoByUser)
		v1.GET("user/:id/passed-audios", api.ListPassedAudioByUser)
		v1.GET("user/:id/comments", api.ListCommentByUser)

		v1.GET("audios/:id", api.ShowAudio)
		v1.GET("audios/:id/comments", api.ListCommentsByMediaId)
		v1.GET("passed-audios", api.ListPassedAudios)

		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		v1.GET("search", api.Search)
		// 其他
		//-------------------用户相关接口------------------------
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user", api.UserRegister)

		// 用户登录
		v1.POST("login", api.UserLogin)

		//验证token，返回用户信息
		v1.GET("verify", api.Verify)
		//刷新token时间
		v1.GET("refresh", api.Refresh)
		//测试token
		v1.GET("sayHello", api.SayHello)

		// 用户登出
		v1.GET("logout", api.UserLogout)

		//获取用户信息
		v1.GET("user/:id", api.GetUser)

		// 需要登录保护的
		authUser := v1.Group("/")
		authUser.Use(middleware.AuthUserRequired())
		{
			//authUser.DELETE("user/:id/logout", api.UserLogout)
			authUser.POST("upload/token", api.MediaUploadToken)
			authUser.POST("videos", api.CreateVideo)
			authUser.PUT("videos/:id", api.UpdateVideo)
			authUser.DELETE("videos/:id", api.DeleteMyVideo)
			authUser.POST("videos/:id/comments", api.PostComment)

			authUser.POST("audios", api.CreateAudio)
			authUser.PUT("audios/:id", api.UpdateAudio)
			authUser.DELETE("audios/:id", api.DeleteAudio)
			authUser.POST("audios/:id/comments", api.PostComment)

			authUser.PUT("user/:id/avatar", api.ChangeUserAvatar)
		}
		// 需要验证审查员身份的
		authInsp := v1.Group("/")
		authInsp.Use(middleware.AuthInspectorRequired())
		{
			authInsp.GET("not-passed-videos", api.ListNotPassedVideos)
			authInsp.GET("not-passed-audios", api.ListNotPassedAudios)
			authInsp.PUT("review/videos/:id", api.DoVideoReview)
			authInsp.PUT("review/audios/:id", api.DoAudioReview)
			authInsp.GET("review/videos/:id", api.GetVideo)
			authInsp.GET("review/audios/:id", api.GetAudio)
			authInsp.GET("videos/:id/sprite", api.GetVideoSpritePic)
			authInsp.GET("videos", api.ListAllVideo)
			authInsp.GET("audios", api.ListAllAudio)
		}
		// 需要验证管理员身份的
		authAdmin := v1.Group("/")
		authAdmin.Use(middleware.AuthAdminRequired())
		{
			authAdmin.PUT("user/:id", api.ChangeUserInfo)
			authAdmin.DELETE("user/:id", api.DeleteUser)
			authAdmin.GET("users", api.ListUser)
		}
	}

	// swagger文档
	// 游览器打开 http://localhost:3000/swagger/index.html
	r.StaticFile("/swagger.json", "./swagger/swagger.json")
	r.Static("/swagger", "./swagger/dist")

	return r
}
