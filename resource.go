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

func (r *Resource) CheckResourceExistsById(c *gin.Context) (uint, error) {
	paramId := c.Param("id")
	idInt64, err := strconv.ParseInt(paramId, 10, 16)
	id := uint(idInt64)
	if err != nil {
		HandleStatusBadRequestError(c, err)
		return id, err
	}

	if !Conn.IsResourceExistsById(r.TableName, idInt64) {
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

func (r *Resource) CreateResource(c *gin.Context, modelPtr interface{}, GuidFiledJsonTag string) {
	if !lightweight_db.IsPointer(modelPtr) {
		HandleInternalServerError(c, errors.New("modelPtr is not type of pointer"))
		return
	}
	if err := c.ShouldBindJSON(modelPtr); err != nil {
		HandleStatusBadRequestError(c, err)
		return
	}

	if GuidFiledJsonTag != "" {
		guidFiled, find := lightweight_db.FieldByJsonTag(lightweight_db.Reflect(modelPtr), GuidFiledJsonTag)
		if !find {
			HandleInternalServerError(c, errors.New("cannot find field: "+GuidFiledJsonTag))
			return
		}
		guidValue := guidFiled.Interface()
		//guidValue := reflect.ValueOf(modelPtr).Elem().FieldByName("Name").String()
		exist, err := r.CheckResourceExistsByGuid(c, GuidFiledJsonTag, guidValue)
		if err != nil {
			HandleInternalServerError(c, err)
			return
		}
		if exist {
			HandleStatusBadRequestError(c, errors.New(fmt.Sprintf("guidField: %s already exists", GuidFiledJsonTag)))
			return
		}
	}

	id, err := Conn.CreateObject(r.TableName, modelPtr)
	if err != nil {
		HandleInternalServerError(c, err)
		return
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

	if err := Conn.DeleteObjectById(r.TableName, int64(id)); err != nil {
		HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
}
