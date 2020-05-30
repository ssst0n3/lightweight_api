package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ssst0n3/awesome_libs"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		Logger.Errorf("error: %+v\n", errors.Errorf(err.Error()))
	}
}

func HandleInternalServerError(c *gin.Context, err error) {
	awesome_libs.CheckErr(err)
	if !c.Writer.Written() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
	}
}

func HandleStatusBadRequestError(c *gin.Context, err error) {
	awesome_libs.CheckErr(err)
	if !c.Writer.Written(){
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  err.Error(),
		})
	}
}
