package lightweight_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"net/http"
	"strconv"
)

func (r *Resource) CheckResourceExistsById(c *gin.Context) (bool, int64, error) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 16)
	if err != nil {
		awesomeError.CheckErr(err)
		return false, id, err
	}

	return Conn.IsResourceExistsById(r.TableName, id), id, err
}

func (r *Resource) MustResourceExistsById(c *gin.Context) (int64, error) {
	exists, id, err := r.CheckResourceExistsById(c)
	if err != nil {
		HandleInternalServerError(c, err)
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

func (r *Resource) CheckResourceExistsByGuid(guidColName string, guidValue interface{}) (bool, error) {
	return Conn.IsResourceExistsByGuid(r.TableName, guidColName, guidValue)
}

func (r *Resource) MustResourceNotExistsByGuid(c *gin.Context, guidColName string, guidValue interface{}) (bool, error) {
	if exists, err := r.CheckResourceExistsByGuid(guidColName, guidValue); err != nil {
		HandleInternalServerError(c, err)
		return false, err
	} else if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  fmt.Sprintf(ResourceAlreadyExists, r.Name, guidColName, guidValue),
		})
		return true, nil
	}
	return false, nil
}













