package middleware

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId  uint `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}
