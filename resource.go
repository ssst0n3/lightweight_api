package lightweight_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_db"
	"net/http"
	"strconv"
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

func (r *Resource) CheckResourceExistsById(c *gin.Context) (int64, error) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 16)
	if err != nil {
		HandleStatusBadRequestError(c, err)
		return id, err
	}

	if !Conn.IsResourceExistsById(r.TableName, id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"reason":  fmt.Sprintf("%sId not exists.", r.Name),
		})
		return id, err
	}
	return id, error(nil)
}

func (r *Resource) CheckResourceExistsByGuid(c *gin.Context, guidColName string, guidValue interface{}) (bool, error) {
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
	lightweight_db.MustIsPointer(modelPtr)
	if GuidFieldJsonTag != "" {
		guidFiled, find := lightweight_db.FieldByJsonTag(lightweight_db.Reflect(modelPtr), GuidFieldJsonTag)
		if !find {
			err := errors.New("cannot find field: " + GuidFieldJsonTag)
			CheckError(err)
			return err
		}
		guidValue := guidFiled.Interface()
		exist, err := r.CheckResourceExistsByGuid(c, GuidFieldJsonTag, guidValue)
		if err != nil {
			CheckError(err)
			return err
		}
		if exist {
			err := errors.New(fmt.Sprintf("guidField: %s already exists", GuidFieldJsonTag))
			CheckError(err)
			return err
		}
	}

	return nil
}

func (r *Resource) MustResourceNotExistsExceptSelfByModelPtr(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string, id int64) error {
	lightweight_db.MustIsPointer(modelPtr)
	if GuidFieldJsonTag != "" {
		guidFiled, find := lightweight_db.FieldByJsonTag(lightweight_db.Reflect(modelPtr), GuidFieldJsonTag)
		if !find {
			err := errors.New("cannot find field: " + GuidFieldJsonTag)
			CheckError(err)
			return err
		}
		guidValue := guidFiled.Interface()
		exist, err := r.CheckResourceExistsByGuidExceptSelf(c, GuidFieldJsonTag, guidValue, id)
		if err != nil {
			CheckError(err)
			return err
		}
		if exist {
			err := errors.New(fmt.Sprintf("guidField: %s already exists", GuidFieldJsonTag))
			CheckError(err)
			return err
		}
	}

	return nil
}

func (r *Resource) CreateResource(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string, taskBeforeCreateObject func(modelPtr interface{})) {
	lightweight_db.MustIsPointer(modelPtr)
	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}

	if err := r.MustResourceNotExistsByModelPtr(c, modelPtr, GuidFieldJsonTag); err != nil {
		HandleInternalServerError(c, err)
		return
	}
	id, err := Conn.CreateObject(r.TableName, modelPtr)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}

	if taskBeforeCreateObject != nil {
		taskBeforeCreateObject(modelPtr)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true, "id": id,
	})
}

func (r *Resource) DeleteResource(c *gin.Context) {
	id, err := r.CheckResourceExistsById(c)
	if err != nil {
		return
	}

	if err := Conn.DeleteObjectById(r.TableName, id); err != nil {
		HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
}

func (r *Resource) UpdateResource(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string) {
	id, err := r.CheckResourceExistsById(c)
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
