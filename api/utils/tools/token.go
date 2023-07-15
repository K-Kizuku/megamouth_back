package tools

import (
	"megamouth/api/entity/models"
	"megamouth/api/usecase/schema"
	"megamouth/api/utils/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(u *models.User) (jwtInfo schema.AuthOutput, err error) {

	token, err := generateToken(u.ID)
	if err != nil {
		return
	}
	jwtInfo = schema.AuthOutput{Jwt: token}

	return
}

func generateToken(userID string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iat": now.Unix(),
		"exp": now.Add(7 * 24 * time.Hour).Unix(),
	})
	return token.SignedString(config.SigningKey)
}
