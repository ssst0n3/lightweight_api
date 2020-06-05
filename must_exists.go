package lightweight_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"net/http"
)

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

func (r *Resource) MustResourceNotExistsByModelPtrWithGuid(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string) error {
	exists, err := r.CheckResourceExistsByModelPtrWithGuid(modelPtr, GuidFieldJsonTag)
	if err != nil {
		HandleInternalServerError(c, err)
		return err
	}
	if exists {
		err := errors.New(fmt.Sprintf(GuidFieldMustNotExists, GuidFieldJsonTag))
		HandleInternalServerError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsExceptSelfByGuid(c *gin.Context, guidColName string, guidValue interface{}, id int64) (bool, error) {
	if exists, err := r.CheckResourceExistsExceptSelfByGuid(guidColName, guidValue, id); err != nil {
		HandleInternalServerError(c, err)
		return false, err
	} else if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  fmt.Sprintf(ResourceMustNotExistsExceptSelf, r.Name),
		})
		return true, nil
	}
	return false, nil
}
