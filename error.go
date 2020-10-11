package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func HandleInternalServerError(c *gin.Context, err error) {
	awesome_error.CheckErr(err)
	if !c.Writer.Written() {
		Response500InternalError(c, err.Error())
	}
}

func HandleStatusBadRequestError(c *gin.Context, err error) {
	awesome_error.CheckErr(err)
	if !c.Writer.Written() {
		Response400BadRequest(c, err.Error())
	}
}
