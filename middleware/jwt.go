package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"net/http"
)

var CloseJwt = false

func GetToken(c *gin.Context) (string, error) {
	token, err := c.Cookie("token")
	if err != nil {
		token = c.GetHeader("token")
	}
	if token == "" {
		return token, errors.New("token empty")
	}
	return token, nil
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !CloseJwt {
			token, err := GetToken(c)
			if err != nil {
				awesome_error.CheckErr(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"auth":   false,
					"reason": "none token",
				})
				return
			}

			claims, err := ParseToken(token)
			if err != nil {
				awesome_error.CheckErr(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"auth":   false,
					"reason": err.Error(),
				})
				return
			}
			if !claims.IsAdmin {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"auth":   false,
					"reason": "you are not admin",
				})
				return
			}
		}
		c.Next()
	}
}

func JwtUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !CloseJwt {
			token, err := c.Cookie("token")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"auth":   false,
					"reason": "none token",
				})
				return
			}

			_, err = ParseToken(token)
			if err != nil {
				awesome_error.CheckErr(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"auth":   false,
					"reason": err.Error(),
				})
				return
			}
		}
		c.Next()
	}
}
