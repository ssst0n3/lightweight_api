package lightweight_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"github.com/ssst0n3/awesome_libs/reflect"
	"net/http"
)

type Resource struct {
	Name             string
	TableName        string
	BaseRelativePath string
	Model            interface{} // model cannot be reused
}

func (r *Resource) ListResource(c *gin.Context) {
	objects, err := Conn.ListAllPropertiesByTableName(r.TableName)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, objects)
}

func (r *Resource) MustResourceNotExistsByGuid(c *gin.Context, guidColName string, guidValue interface{}) (bool, error) {
	if exists, err := Conn.IsResourceExistsByGuid(r.TableName, guidColName, guidValue); err != nil {
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

func (r *Resource) CheckResourceExistsByGuidExceptSelf(c *gin.Context, guidColName string, guidValue interface{}, id int64) (bool, error) {
	if exists, err := Conn.IsResourceExistsExceptSelfByGuid(r.TableName, guidColName, guidValue, id); err != nil {
		HandleInternalServerError(c, err)
		return false, err
	} else if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  fmt.Sprintf("%s name already exists.", r.Name),
		})
		return true, nil
	}
	return false, nil
}

func (r *Resource) MustResourceNotExistsByModelPtr(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string) error {
	reflect.MustPointer(modelPtr)
	if GuidFieldJsonTag != "" {
		guidFiled, find := reflect.FieldByJsonTag(reflect.Value(modelPtr), GuidFieldJsonTag)
		if !find {
			err := errors.New("cannot find field: " + GuidFieldJsonTag)
			awesomeError.CheckErr(err)
			return err
		}
		guidValue := guidFiled.Interface()
		exist, err := r.MustResourceNotExistsByGuid(c, GuidFieldJsonTag, guidValue)
		if err != nil {
			awesomeError.CheckErr(err)
			return err
		}
		if exist {
			err := errors.New(fmt.Sprintf("guidField: %s already exists", GuidFieldJsonTag))
			awesomeError.CheckErr(err)
			return err
		}
	}

	return nil
}

func (r *Resource) MustResourceNotExistsExceptSelfByModelPtr(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string, id int64) error {
	reflect.MustPointer(modelPtr)
	if GuidFieldJsonTag != "" {
		guidFiled, find := reflect.FieldByJsonTag(reflect.Value(modelPtr), GuidFieldJsonTag)
		if !find {
			err := errors.New("cannot find field: " + GuidFieldJsonTag)
			awesomeError.CheckErr(err)
			return err
		}
		guidValue := guidFiled.Interface()
		exist, err := r.CheckResourceExistsByGuidExceptSelf(c, GuidFieldJsonTag, guidValue, id)
		if err != nil {
			awesomeError.CheckErr(err)
			return err
		}
		if exist {
			err := errors.New(fmt.Sprintf("guidField: %s already exists", GuidFieldJsonTag))
			awesomeError.CheckErr(err)
			return err
		}
	}

	return nil
}

func (r *Resource) CreateResource(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string, taskBeforeCreateObject func(modelPtr interface{})) {
	reflect.MustPointer(modelPtr)
	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}

	if err := r.MustResourceNotExistsByModelPtr(c, modelPtr, GuidFieldJsonTag); err != nil {
		HandleInternalServerError(c, err)
		return
	}
	if taskBeforeCreateObject != nil {
		taskBeforeCreateObject(modelPtr)
	}
	id, err := Conn.CreateObject(r.TableName, modelPtr)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}

	Response200CreateSuccess(c, uint(id))
}

func (r *Resource) DeleteResource(c *gin.Context) {
	id, err := r.MustResourceExistsById(c)
	if err != nil {
		return
	}

	if err := Conn.DeleteObjectById(r.TableName, id); err != nil {
		HandleInternalServerError(c, err)
	} else {
		Response200DeleteSuccess(c)
	}
}

func (r *Resource) UpdateResource(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string, taskBeforeCreateObject func(modelPtr interface{})) {
	id, err := r.MustResourceExistsById(c)
	if err != nil {
		return
	}

	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}

	if err := r.MustResourceNotExistsExceptSelfByModelPtr(c, modelPtr, GuidFieldJsonTag, id); err != nil {
		HandleInternalServerError(c, err)
		return
	}
	if taskBeforeCreateObject != nil {
		taskBeforeCreateObject(modelPtr)
	}
	if err := Conn.UpdateObject(id, r.TableName, modelPtr); err != nil {
		HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
}

func (r *Resource) ShowResource(c *gin.Context) {
	id, err := r.MustResourceExistsById(c)
	if err != nil {
		return
	}
	result, err := Conn.OrmShowObjectByIdUsingReflectRet(r.TableName, id, r.Model)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
