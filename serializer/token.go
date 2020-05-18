package serializer

import (
	"giligili/token"
)

type LoginResponse struct {
	Response
	Data string           `json:"data"`
	User *token.JWTClaims `json:"user"`
}

type ClaimResponse struct {
	Response
	Data token.JWTClaims `json:"data"`
}

func BuildLoginResponse(tokenStr string) LoginResponse {
	claims, _ := token.VerifyAction(tokenStr)
	return LoginResponse{
		Data: tokenStr,
		User: claims,
	}
}

func BuildClaimResponse(claim token.JWTClaims) ClaimResponse {
	return ClaimResponse{
		Data: claim,
	}
}
