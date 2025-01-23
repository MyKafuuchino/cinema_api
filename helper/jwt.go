package helper

import (
	"cinema_api/config"
	"cinema_api/types"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWTToken(payload types.UserPayload) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	secretKey := []byte(config.GlobalAppConfig.SecretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "my-auth-server",
		"id":   payload.Id,
		"role": payload.Role,
		"exp":  expirationTime,
	})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func DecodeJWTToken(tokenString string) (types.UserPayload, error) {
	secretKey := []byte(config.GlobalAppConfig.SecretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return types.UserPayload{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payload := types.UserPayload{
			Id:   uint(claims["id"].(float64)),
			Role: claims["role"].(string),
		}

		return payload, nil
	}

	return types.UserPayload{}, errors.New("invalid token")
}
