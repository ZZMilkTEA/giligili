package token

import (
	"errors"
	"fmt"
	"giligili/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
	ErrorReason_TimesOut   = "登陆超时"
)

var (
	Secret     = "zzmilktea" // 加盐,实际生产环境需要更改
	ExpireTime = 3600        // token有效期
)

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Permissions uint   `json:"permissions"`
}

func CreateUserToken(user model.User) (tokenString string, err error) {
	claims := &JWTClaims{
		UserID:      user.ID,
		Username:    user.UserName,
		Nickname:    user.Nickname,
		Permissions: 0,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	tokenString, err = GetToken(claims)
	return
}

//注：签名方式实际生产环境是需要保密的
func GetToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}

func VerifyAction(strToken string) (*JWTClaims, error) {

	var fun = func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	}

	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, fun)
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}

	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify success,token:" + strToken)
	return claims, nil
}

func GetLoggedUserId(strToken string) (userId uint, err error) {
	claim, err := VerifyAction(strToken)
	return claim.UserID, err
}
