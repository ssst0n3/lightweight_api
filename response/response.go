package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success200(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Base{
		Success: true,
		Message: msg,
	})
}

func CreateSuccess200(c *gin.Context, id uint, msg string) {
	c.JSON(http.StatusOK, CreateSuccess{
		Base: Base{
			Success: true,
			Message: msg,
		},
		Id: id,
	})
}

func Error(c *gin.Context, statusCode int, reason string) {
	c.JSON(statusCode, Err{
		Base: Base{
			Success: false,
		},
		Reason: reason,
	})
}

func InternalError500(c *gin.Context, reason string) {
	Error(c, http.StatusInternalServerError, reason)
}

func BadRequest400(c *gin.Context, reason string) {
	Error(c, http.StatusBadRequest, reason)
}

func Unauthorized401(c *gin.Context, reason string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized,
		Auth{
			Base: Base{
				Success: false,
				Message: reason,
			},
			Auth: false,
		},
	)
}

func UpdateSuccess200(c *gin.Context) {
	Success200(c, "update success")
}

func DeleteSuccess200(c *gin.Context) {
	Success200(c, "delete success")
}
