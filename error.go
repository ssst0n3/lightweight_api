package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		Logger.Errorf("error: %+v\n", errors.Errorf(err.Error()))
	}
}

func HandleInternalServerError(c *gin.Context, err error) {
	CheckError(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"reason":  err.Error(),
	})
}

func HandleStatusBadRequestError(c *gin.Context, err error) {
	CheckError(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"reason":  err.Error(),
	})
}
