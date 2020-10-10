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
	GuidFieldJsonTag string
}

func (r *Resource) ListResource(c *gin.Context) {
	objects, err := Conn.ListAllPropertiesByTableName(r.TableName)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, objects)
}

func (r *Resource) MapResourceById(c *gin.Context) {
	objects, err := Conn.MapAllPropertiesByTableName(r.TableName)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, objects)
}

func (r *Resource) CreateResourceTemplate(c *gin.Context, taskBeforeCreateObject func(modelPtr interface{}) error, taskAfterCreateObject func(id int64) error) {
	awesome_reflect.MustNotPointer(r.Model)
	modelPtr := awesome_reflect.EmptyPointerOfModel(r.Model)
	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}
	if len(r.GuidFieldJsonTag) > 0 {
		if err := r.MustResourceNotExistsByModelPtrWithGuid(c, modelPtr, r.GuidFieldJsonTag); err != nil {
			return
		}
	}
	if taskBeforeCreateObject != nil {
		if err := taskBeforeCreateObject(modelPtr); err != nil {
			HandleInternalServerError(c, err)
			return
		}
	}
	id, err := Conn.CreateObject(r.TableName, modelPtr)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
	if taskAfterCreateObject != nil {
		if err := taskAfterCreateObject(id); err != nil {
			HandleInternalServerError(c, err)
			return
		}
	}
	Response200CreateSuccess(c, uint(id))
}

func (r *Resource) CreateResource(c *gin.Context) {
	r.CreateResourceTemplate(c, nil, nil)
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

func (r *Resource) UpdateResourceTemplate(c *gin.Context, model interface{}, taskBeforeCreateObject func(modelPtr interface{}) error) {
	if model == nil {
		model = r.Model
	}
	awesome_reflect.MustNotPointer(model)
	modelPtr := awesome_reflect.EmptyPointerOfModel(model)
	id, err := r.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}

	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}

	if len(r.GuidFieldJsonTag) > 0 {
		if err := r.MustResourceNotExistsExceptSelfByModelPtrWithGuid(c, modelPtr, r.GuidFieldJsonTag, id); err != nil {
			HandleInternalServerError(c, err)
			return
		}
	}
	if taskBeforeCreateObject != nil {
		if err := taskBeforeCreateObject(modelPtr); err != nil {
			HandleInternalServerError(c, err)
			return
		}
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

func (r *Resource) UpdateResource(c *gin.Context) {
	r.UpdateResourceTemplate(c, nil, nil)
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
