package model

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type HandleRefreshJWTRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type HandleRefreshJWTResponse struct {
	NewAccessToken string `json:"new_access_token" binding:"required"`
}
