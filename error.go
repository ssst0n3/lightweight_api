package lightweight_api

import (
	"github.com/gin-gonic/gin"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"net/http"
)

func HandleInternalServerError(c *gin.Context, err error) {
	awesomeError.CheckErr(err)
	if !c.Writer.Written() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
	}
}

func HandleStatusBadRequestError(c *gin.Context, err error) {
	awesomeError.CheckErr(err)
	if !c.Writer.Written() {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
	}
}
