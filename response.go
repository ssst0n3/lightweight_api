package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response200CreateSuccess(c *gin.Context, id uint) {
	c.JSON(http.StatusOK, gin.H{
		"success": true, "id": id,
	})
}

func Response200Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func Response200UpdateSuccess(c *gin.Context)  {
	Response200Success(c)
}

func Response200DeleteSuccess(c *gin.Context)  {
	Response200Success(c)
}