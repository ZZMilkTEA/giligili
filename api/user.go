package api

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/userService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service userService.UserRegisterService
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
	service := userService.ShowUserService{}
	user := service.Show(c.Param("id"))
	c.JSON(200, user)
}

// ListUser 所有用户详情
func ListUser(c *gin.Context) {
	service := userService.ListUserService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ChangeUserPermission 改变用户权限 (未来将删除
func ChangeUserPermission(c *gin.Context) {
	service := userService.ChangeUserPermissionService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ChangeUserPermission()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ChangeUserInfo 改变用户信息
func ChangeUserInfo(c *gin.Context) {
	userIdStr := c.Param("id")
	temp, _ := strconv.ParseUint(userIdStr, 10, 32)
	userId := uint(temp)

	service := userService.ChangeUserInfoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ChangeUserInfo(userId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ChangeUserAvatar 改变用户头像
func ChangeUserAvatar(c *gin.Context) {
	userIdStr := c.Param("id")
	temp, _ := strconv.ParseUint(userIdStr, 10, 32)
	userId := uint(temp)

	userStr, _ := c.Get("user")
	user, _ := userStr.(*model.User)

	if user.ID != userId {
		errResponse := serializer.Response{
			Status: 40003,
			Data:   nil,
			Msg:    "只能修改自己的头像",
			Error:  "verify err",
		}
		c.JSON(200, errResponse)
	}

	service := userService.ChangeUserInfoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ChangeUserAvatar(user.ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteUser 注销用户
func DeleteUser(c *gin.Context) {
	service := userService.DeleteUserService{}
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
