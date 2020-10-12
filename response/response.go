package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSuccess200(c *gin.Context, id uint, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"success": true, "id": id, "msg": msg,
	})
}

func Success200(c *gin.Context, msg string, h gin.H) {
	resp := gin.H{
		"success": true,
		"msg":     msg,
	}
	for k, v := range h {
		resp[k] = v
	}
	c.JSON(http.StatusOK, resp)
}

func Error(c *gin.Context, statusCode int, reason string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"reason":  reason,
	})
}

func InternalError500(c *gin.Context, reason string) {
	Error(c, http.StatusInternalServerError, reason)
}

func BadRequest400(c *gin.Context, reason string) {
	Error(c, http.StatusBadRequest, reason)
}

func Unauthorized401(c *gin.Context, reason string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"auth":    false,
		"reason":  reason,
	})
}

func UpdateSuccess200(c *gin.Context) {
	Success200(c, "update success", nil)
}

func DeleteSuccess200(c *gin.Context) {
	Success200(c, "delete success", nil)
}
