package jwtToken

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessJWT(username, id, accessSecretKey string, expiryMinutes int) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_x_username": username,
		"_x_user_id":  id,
		"exp":         time.Now().Add(time.Minute * time.Duration(expiryMinutes)).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(accessSecretKey))
	if err != nil {
		return "Unsuccessfull", err
	}

	return accessTokenString, nil
}

func GenerateRefreshJWT(username, id, refreshSecretKey string, expiryHours int) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_x_data_username": username,
		"_x_data_id":       id,
		"exp":              time.Now().Add(time.Hour * time.Duration(expiryHours)).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		return "Unsuccessfull", err
	}

	return refreshTokenString, nil
}
