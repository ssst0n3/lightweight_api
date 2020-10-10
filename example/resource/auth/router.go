package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api/middleware"
)

func InitRouter(router *gin.Engine) {
	v1AuthGroup := router.Group(AuthResource.BaseRelativePath)
	v1AuthGroupAdmin := router.Group(AuthResource.BaseRelativePath, middleware.JwtAdmin())
	{
		v1AuthGroup.POST("", Login)
		v1AuthGroupAdmin.GET("", func(context *gin.Context) {})
	}
}