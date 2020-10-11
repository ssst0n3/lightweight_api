package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api/response"
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
				response.Unauthorized401(c, "none token.")
				return
			}

			claims, err := ParseToken(token)
			if err != nil {
				awesome_error.CheckErr(err)
				response.Unauthorized401(c, err.Error())
				return
			}
			if !claims.IsAdmin {
				response.Unauthorized401(c, "you are not admin.")
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
				response.Unauthorized401(c, "none token.")
				return
			}

			_, err = ParseToken(token)
			if err != nil {
				awesome_error.CheckErr(err)
				response.Unauthorized401(c, err.Error())
				return
			}
		}
		c.Next()
	}
}
