package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api/middleware"
)

func InitRouter(router *gin.Engine) {
	anonymous := router.Group(Resource.BaseRelativePath)
	admin := router.Group(Resource.BaseRelativePath, middleware.JwtAdmin())
	{
		anonymous.POST("", Login)
		admin.GET("", Nothing)
	}
}
