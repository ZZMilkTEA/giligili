package api

import (
	"giligili/serializer"
	"giligili/service"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetUser 用户详情
func GetUser(c *gin.Context) {
	service := service.ShowUserService{}
	user := service.Show(c.Param("id"))
	c.JSON(200, user)
}

// ListUser 所有用户详情
func ListUser(c *gin.Context) {
	service := service.ListUserService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ChangeUserPermission 改变用户权限
func ChangeUserPermission(c *gin.Context) {
	service := service.ChangeUserPermissionService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ChangeUserPermission()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteUser 注销用户
func DeleteUser(c *gin.Context) {
	service := service.DeleteUserService{}
	user := service.Delete(c.Param("id"))
	c.JSON(200, user)
}

//// UserLogin 用户登录接口 (session)
//func UserLogin(c *gin.Context) {
//	var service service.UserLoginService
//	if err := c.ShouldBind(&service); err == nil {
//		if user, err := service.Login(); err != nil {
//			c.SetCookie("user_id", strconv.FormatUint(uint64(user.ID), 10), 3600, "/", "localhost", false, true)
//			c.JSON(200, err)
//		} else {
//			// 设置Session
//			s := sessions.Default(c)
//			s.Clear()
//			s.Set("user_id", user.ID)
//			s.Save()
//
//			res := serializer.BuildUserResponse(user)
//			c.JSON(200, res)
//		}
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
