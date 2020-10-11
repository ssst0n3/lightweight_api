package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response200CreateSuccess(c *gin.Context, id uint, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"success": true, "id": id, "msg": msg,
	})
}

func Response200Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     msg,
	})
}

func ResponseError(c *gin.Context, statusCode int, reason string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"reason":  reason,
	})
}

func Response500InternalError(c *gin.Context, reason string) {
	ResponseError(c, http.StatusInternalServerError, reason)
}

func Response400BadRequest(c *gin.Context, reason string)  {
	ResponseError(c, http.StatusBadRequest, reason)
}

func Response200UpdateSuccess(c *gin.Context) {
	Response200Success(c, "update success")
}

func Response200DeleteSuccess(c *gin.Context) {
	Response200Success(c, "delete success")
}
