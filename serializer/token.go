package serializer

import (
	"giligili/token"
)

type TokenResponse struct {
	Response
	Data string `json:"data"`
}

type ClaimResponse struct {
	Response
	Data token.JWTClaims `json:"data"`
}

func BuildTokenResponse(token string) TokenResponse {
	return TokenResponse{
		Data: token,
	}
}

func BuildClaimResponse(claim token.JWTClaims) ClaimResponse {
	return ClaimResponse{
		Data: claim,
	}
}
