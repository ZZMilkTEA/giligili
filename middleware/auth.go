package middleware

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/token"
	"github.com/gin-gonic/gin"
)

//CurrentUser 获取登录用户(从token)
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("token")
		if tokenStr == "" || tokenStr == "[object Object]" {
			c.Next()
		} else {
			claim, err := token.VerifyAction(tokenStr)
			if err != nil {
				c.Next()
			} else {
				uid := claim.UserID
				user, err := model.GetUser(uid)
				if err == nil {
					c.Set("user", &user)
				}
				c.Next()
			}
		}
	}
}

// AuthUserRequired 需要登录
func AuthUserRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.Response{
			Status: 401,
			Msg:    "登陆验证不通过,请重新登陆",
		})
		c.Abort()
	}
}

// AuthInspectorRequired 需要审核员登录
func AuthInspectorRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if userModel, ok := user.(*model.User); ok && userModel.Permission >= 1 {
				c.Next()
				return
			} else {
				c.JSON(200, serializer.Response{
					Status: 401,
					Msg:    "需要更高的权限",
				})
			}
		} else {
			c.JSON(200, serializer.Response{
				Status: 401,
				Msg:    "需要登陆",
			})
		}
		c.Abort()
	}
}
