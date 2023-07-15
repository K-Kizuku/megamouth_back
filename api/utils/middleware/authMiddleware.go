package middleware

import (
	"fmt"
	"megamouth/api/utils/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *gin.Context) {
	AuthorizationHeader := c.Request.Header.Get("Authorization")
	if AuthorizationHeader != "" {
		ary := strings.Split(AuthorizationHeader, " ")
		if len(ary) == 2 {
			if ary[0] == "Bearer" {
				t, err := jwt.ParseWithClaims(ary[1], &jwt.MapClaims{}, func(token *jwt.Token) (any, error) {
					return config.SigningKey, nil
				})
				if claims, ok := t.Claims.(*jwt.MapClaims); ok && t.Valid {
					userId := (*claims)["sub"].(string)
					c.Set("user_id", userId)
				} else {
					fmt.Println(err)
				}
			}
		}
	}
	c.Next()
}
