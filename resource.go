package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resource struct {
	Name             string
	TableName        string
	BaseRelativePath string
	Model            interface{}
}

func (r *Resource) ListResource(c *gin.Context) {
	objects, err := Conn.ListAllPropertiesByTableName(r.TableName)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, objects)
}
