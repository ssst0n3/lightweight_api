package lightweight_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"github.com/ssst0n3/lightweight_api/db"
	"github.com/ssst0n3/lightweight_api/response"
	"net/http"
	"reflect"
	"strconv"
)

type Resource struct {
	Name             string
	TableName        string
	BaseRelativePath string
	Model            interface{} // model cannot be reused
	GuidFieldJsonTag string
}

func NewResource(resourceName string, tableName string, model interface{}, guidFiledJsonTag string) Resource {
	awesome_reflect.MustNotPointer(model)
	return Resource{
		Name:             resourceName,
		TableName:        tableName,
		BaseRelativePath: BaseRelativePathV1(resourceName),
		Model:            model,
		GuidFieldJsonTag: guidFiledJsonTag,
	}
}

func BaseRelativePath(apiVersion, resourceName string) string {
	return fmt.Sprintf("/api/%s/%s", apiVersion, resourceName)
}

func BaseRelativePathV1(resourceName string) string {
	return BaseRelativePath("v1", resourceName)
}

func (r *Resource) CountResource(c *gin.Context) {
	var count int64
	model := awesome_reflect.EmptyPointerOfModel(r.Model)
	DB.Model(model).Count(&count)
	c.String(http.StatusOK, strconv.Itoa(int(count)))
}

func (r *Resource) ListResource(c *gin.Context) {
	var objects []map[string]interface{}
	model := awesome_reflect.EmptyPointerOfModel(r.Model)
	DB.Table(r.TableName).Model(model).Find(&objects)
	if objects == nil {
		objects = []map[string]interface{}{}
	}
	c.JSON(http.StatusOK, objects)
}

func (r *Resource) MapResourceById(c *gin.Context) {
	var list []map[string]interface{}
	model := awesome_reflect.EmptyPointerOfModel(r.Model)
	DB.Model(model).Find(&list)
	objects := db.MapObjectById(list)
	c.JSON(http.StatusOK, objects)
}

func (r *Resource) CreateResourceTemplate(c *gin.Context, taskBeforeCreateObject func(modelPtr interface{}) error, taskAfterCreateObject func(id uint) error) {
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

	DB.Create(modelPtr)
	id := awesome_reflect.ValueByPtr(modelPtr).FieldByName("ID").Interface().(uint)

	if taskAfterCreateObject != nil {
		if err := taskAfterCreateObject(id); err != nil {
			HandleInternalServerError(c, err)
			return
		}
	}
	var msg string
	if len(r.GuidFieldJsonTag) > 0 {
		guidValue, _ := awesome_reflect.FieldByJsonTag(reflect.ValueOf(modelPtr).Elem(), r.GuidFieldJsonTag)
		msg = fmt.Sprintf(MsgResourceCreateSuccess, r.Name, guidValue)
	} else {
		msg = fmt.Sprintf(MsgResourceCreateSuccess, r.Name, id)
	}
	if !c.IsAborted() {
		response.CreateSuccess200(c, id, msg)
	}
}

func (r *Resource) CreateResource(c *gin.Context) {
	r.CreateResourceTemplate(c, nil, nil)
}

func (r *Resource) DeleteResource(c *gin.Context) {
	r.DeleteResourceTemplate(c, false)
}

func (r *Resource) DeleteResourceTemplate(c *gin.Context, softDelete bool) {
	id, err := r.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}

	model := awesome_reflect.EmptyPointerOfModel(r.Model)
	if softDelete {
		DB.Delete(model, id) // soft delete
	} else {
		DB.Unscoped().Delete(model, id)
	}
	response.DeleteSuccess200(c)
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
		if err := r.MustResourceNotExistsExceptSelfByModelPtrWithGuid(c, modelPtr, r.GuidFieldJsonTag, uint(id)); err != nil {
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

	awesome_reflect.ValueByPtr(modelPtr).FieldByName("ID").SetUint(uint64(id))
	result := DB.Updates(modelPtr)
	if result.Error != nil {
		HandleInternalServerError(c, err)
		return
	} else {
		response.UpdateSuccess200(c)
		return
	}
	//if err := Conn.UpdateObject(id, r.TableName, modelPtr); err != nil {
	//	HandleInternalServerError(c, err)
	//	return
	//} else {
	//	response.UpdateSuccess200(c)
	//	return
	//}
}

func (r *Resource) UpdateResource(c *gin.Context) {
	r.UpdateResourceTemplate(c, nil, nil)
}

func (r *Resource) ShowResource(c *gin.Context) {
	id, err := r.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}

	model := awesome_reflect.EmptyPointerOfModel(r.Model)
	result := DB.Table(r.TableName).First(model, id)
	if result.Error != nil {
		HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, model)
}
