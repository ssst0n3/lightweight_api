package lightweight_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"strconv"
)

func (r *Resource) CheckResourceExistsByIdAutoParseParam(c *gin.Context) (bool, int64, error) {
	paramId := c.Param("id")
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		awesome_error.CheckErr(err)
		return false, id, err
	}
	exists, err := r.CheckResourceExistsById(uint(id))
	return exists, id, err
}

func (r *Resource) CheckResourceExistsById(id uint) (exists bool, err error) {
	var count int64
	model := awesome_reflect.EmptyPointerOfModel(r.Model)
	err = DB.Find(model, id).Count(&count).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	exists = count > 0
	return
}

func (r *Resource) CheckResourceExistsByGuid(guidColName string, guidValue interface{}) (exists bool) {
	var count int64
	err := DB.Table(r.TableName).Where(map[string]interface{}{guidColName: guidValue}).Count(&count).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	exists = count > 0
	return
}

func (r *Resource) CheckResourceExistsByModelPtrWithGuid(modelPtr interface{}, GuidFieldJsonTag string) bool {
	awesome_reflect.MustPointer(modelPtr)
	if GuidFieldJsonTag == "" {
		// please make sure by developer
		panic(GuidTagMustNotBeEmpty)
	}
	guidFiled, find := awesome_reflect.FieldByJsonTag(awesome_reflect.Value(modelPtr), GuidFieldJsonTag)
	if !find {
		// please make sure by developer
		panic(fmt.Sprintf(FieldCannotFind, GuidFieldJsonTag))
	}
	guidValue := guidFiled.Interface()
	exists := r.CheckResourceExistsByGuid(GuidFieldJsonTag, guidValue)
	return exists
}

func (r *Resource) CheckResourceExistsExceptSelfByGuid(guidColName string, guidValue interface{}, id uint) bool {
	var count int64
	whereQuery := awesome_libs.Format("{.guid}=? AND id <>? ", awesome_libs.Dict{"guid": guidColName})
	DB.Table(r.TableName).Where(whereQuery, guidValue, id).Count(&count)
	return count > 0
}

func (r *Resource) CheckResourceExistsExceptSelfByModelPtrWithGuid(modelPtr interface{}, GuidFieldJsonTag string, id uint) bool {
	awesome_reflect.MustPointer(modelPtr)
	if GuidFieldJsonTag == "" {
		// please make sure by developer
		panic(GuidTagMustNotBeEmpty)
	}
	guidFiled, find := awesome_reflect.FieldByJsonTag(awesome_reflect.Value(modelPtr), GuidFieldJsonTag)
	if !find {
		// please make sure by developer
		panic(fmt.Sprintf(FieldCannotFind, GuidFieldJsonTag))
	}
	guidValue := guidFiled.Interface()
	exists := r.CheckResourceExistsExceptSelfByGuid(GuidFieldJsonTag, guidValue, id)
	return exists
}
