package kv_config

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	group := router.Group(Resource.BaseRelativePath)
	{
		group.GET("/", Resource.ListResource)
		group.GET("/:key", Get)
		group.POST("", Resource.CreateResource)
		group.PUT("/:key", Update)
		group.DELETE("/:key", Delete)
	}
}