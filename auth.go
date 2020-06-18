package lightweight_api

import (
	"github.com/gin-gonic/gin"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"github.com/ssst0n3/lightweight_api/middleware"
)

func CheckIsAdmin(c *gin.Context) bool {
	token, err := middleware.GetToken(c)
	if err != nil {
		return false
	}

	claims, err := middleware.ParseToken(token)
	if err != nil {
		return false
	}
	return claims.IsAdmin
}

func GetUserId(c *gin.Context) (uint, error) {
	token, err := middleware.GetToken(c)
	if err != nil {
		return 0, err
	}

	claims, err := middleware.ParseToken(token)
	if err != nil {
		return 0, err
	}
	userId := claims.UserId
	return userId, nil
}

func GetClaims(c *gin.Context) (*middleware.Claims, error) {
	token, err := middleware.GetToken(c)
	if err != nil {
		awesomeError.CheckErr(err)
		return nil, err
	}
	return middleware.ParseToken(token)
}
