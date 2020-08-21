package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
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

func (r *Resource) CreateResource(
	c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string,
	taskBeforeCreateObject func(modelPtr interface{}), taskAfterCreateObject func(id int64),
) {
	awesome_reflect.MustPointer(modelPtr)
	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}

	if len(GuidFieldJsonTag) > 0 {
		if err := r.MustResourceNotExistsByModelPtrWithGuid(c, modelPtr, GuidFieldJsonTag); err != nil {
			return
		}
	}
	if taskBeforeCreateObject != nil {
		taskBeforeCreateObject(modelPtr)
	}
	id, err := Conn.CreateObject(r.TableName, modelPtr)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
	if taskAfterCreateObject != nil {
		taskAfterCreateObject(id)
	}

	Response200CreateSuccess(c, uint(id))
}

func (r *Resource) DeleteResource(c *gin.Context) {
	id, err := r.MustResourceExistsByIdAutoParseParam(c)
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
	id, err := r.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}

	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}

	if len(GuidFieldJsonTag) > 0 {
		if err := r.MustResourceNotExistsExceptSelfByModelPtrWithGuid(c, modelPtr, GuidFieldJsonTag, id); err != nil {
			HandleInternalServerError(c, err)
			return
		}
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
	id, err := r.MustResourceExistsByIdAutoParseParam(c)
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
