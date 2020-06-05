package lightweight_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"github.com/ssst0n3/awesome_libs/reflect"
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

func (r *Resource) CheckResourceExistsByGuid(guidColName string, guidValue interface{}) (bool, error) {
	return Conn.IsResourceExistsByGuid(r.TableName, guidColName, guidValue)
}

func (r *Resource) CheckResourceExistsByModelPtrWithGuid(modelPtr interface{}, GuidFieldJsonTag string) (bool, error) {
	reflect.MustPointer(modelPtr)
	if GuidFieldJsonTag == "" {
		// please make sure by developer
		panic(GuidTagMustNotBeEmpty)
	}
	guidFiled, find := reflect.FieldByJsonTag(reflect.Value(modelPtr), GuidFieldJsonTag)
	if !find {
		// please make sure by developer
		panic(fmt.Sprintf(FieldCannotFind, GuidFieldJsonTag))
	}
	guidValue := guidFiled.Interface()
	exists, err := r.CheckResourceExistsByGuid(GuidFieldJsonTag, guidValue)
	if err != nil {
		awesomeError.CheckErr(err)
		return false, err
	}
	return exists, nil
}

func (r *Resource) CheckResourceExistsExceptSelfByGuid(guidColName string, guidValue interface{}, id int64) (bool, error) {
	return Conn.IsResourceExistsExceptSelfByGuid(r.TableName, guidColName, guidValue, id)
}
