package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api/response"
)

func HandleInternalServerError(c *gin.Context, err error) {
	awesome_error.CheckErr(err)
	if !c.Writer.Written() {
		response.InternalError500(c, err.Error())
	}
}

func HandleStatusBadRequestError(c *gin.Context, err error) {
	awesome_error.CheckErr(err)
	if !c.Writer.Written() {
		response.BadRequest400(c, err.Error())
	}
}
