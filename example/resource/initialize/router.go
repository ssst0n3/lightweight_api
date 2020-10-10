package initialize

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	anonymous := router.Group(Resource.BaseRelativePath)
	{
		anonymous.POST("/create_user", CreateUser)
		anonymous.POST("/end", End)
	}
}
