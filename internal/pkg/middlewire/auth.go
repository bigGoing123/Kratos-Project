package middlewire

import (
	jwt5 "github.com/golang-jwt/jwt/v5"
)

// 生成token
func GenerateToken(secret, username string) string {

	token := jwt5.NewWithClaims(jwt5.SigningMethodHS256, jwt5.MapClaims{
		"username": username,
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
