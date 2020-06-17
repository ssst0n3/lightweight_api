package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"net/http"
)

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

func JwtAdmin(closeJwt bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !closeJwt {
			token, err := GetToken(c)
			if err != nil {
				awesomeError.CheckErr(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"auth":   false,
					"reason": "none token",
				})
				return
			}

			claims, err := ParseToken(token)
			if err != nil {
				awesomeError.CheckErr(err)
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

func JwtUser(closeJwt bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !closeJwt {
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
				awesomeError.CheckErr(err)
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
