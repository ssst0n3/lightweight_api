package lightweight_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (r *Resource) CheckResourceExistsById(c *gin.Context) (bool, int64, error) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 16)
	if err != nil {
		HandleStatusBadRequestError(c, err)
		return false, id, err
	}

	return Conn.IsResourceExistsById(r.TableName, id), id, err
}

func (r *Resource) MustResourceExistsById(c *gin.Context) (int64, error) {
	exists, id, err := r.CheckResourceExistsById(c)
	if err != nil {
		return id, err
	}
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  fmt.Sprintf(ResourceMustExists, r.Name),
		})
	}
	return id, nil
}
