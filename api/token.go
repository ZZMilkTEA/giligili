package api

import (
	"giligili/serializer"
	"giligili/service"
	"giligili/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SayHello(c *gin.Context) {
	strToken, err := c.Cookie("Token")
	claim, err := token.VerifyAction(strToken)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	c.String(200, "hello,", claim.Username)
}

func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Login(); err == nil {
			signedToken, err := token.CreateUserToken(user)
			if err != nil {
				c.JSON(200, err)
				return
			}
			c.SetCookie("Token", signedToken, 3600, "/", "localhost", false, true)
			c.JSON(200, signedToken)
		} else {
			c.JSON(200, err)
			return
		}
	} else {
		c.JSON(http.StatusNotFound, err)
		return
	}
}

func Verify(c *gin.Context) {
	strToken, err := c.Cookie("Token")
	if err != nil {
		c.JSON(200, err)
		return
	}
	claim, err := token.VerifyAction(strToken)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, claim)
}

func Refresh(c *gin.Context) {
	strToken, err := c.Cookie("Token")
	claims, err := token.VerifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := token.GetToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	c.SetCookie("Token", "", 3600, "/", "localhost", false, true)
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}
