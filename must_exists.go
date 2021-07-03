package lightweight_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

func (r Resource) MustResourceExistsGetModelByIdAutoParseParam(c *gin.Context) (id int64, model interface{}, err error) {
	id, err = r.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}
	//model, err = Conn.OrmShowObjectByIdUsingReflectRet(r.TableName, id, r.User)
	//if err != nil {
	//	HandleInternalServerError(c, err)
	//	return
	//}
	//m := awesome_reflect.EmptyPointerOfModel(r.User)
	m := reflect.New(reflect.TypeOf(r.Model))
	DB.Table(r.TableName).First(m.Interface(), id)
	model = m.Elem().Interface()
	return
}

func (r *Resource) MustResourceExistsByIdAutoParseParam(c *gin.Context) (int64, error) {
	exists, id, err := r.CheckResourceExistsByIdAutoParseParam(c)
	if err != nil {
		HandleInternalServerError(c, err)
		return id, err
	}
	if !exists {
		err := errors.New(fmt.Sprintf(ResourceMustExists, r.Name))
		HandleStatusBadRequestError(c, err)
		return id, err
	}
	return id, nil
}

func (r *Resource) MustResourceExistsById(c *gin.Context, id uint) error {
	exists, err := r.CheckResourceExistsById(id)
	if err != nil {
		HandleStatusBadRequestError(c, err)
		return err
	}
	if !exists {
		err := errors.New(fmt.Sprintf(ResourceMustExists, r.Name))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceExistsByGuid(c *gin.Context, guidColName string, guidValue interface{}) error {
	if !r.CheckResourceExistsByGuid(guidColName, guidValue) {
		err := errors.New(fmt.Sprintf(ResourceMustExists, r.Name))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsByGuid(c *gin.Context, guidColName string, guidValue interface{}) error {
	if r.CheckResourceExistsByGuid(guidColName, guidValue) {
		//err := errors.New(fmt.Sprintf(ResourceAlreadyExists, r.Name, guidColName, guidValue))
		err := errors.New(fmt.Sprintf(GuidFieldMustNotExists, guidColName))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsByModelPtrWithGuid(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string) error {
	if r.CheckResourceExistsByModelPtrWithGuid(modelPtr, GuidFieldJsonTag) {
		err := errors.New(fmt.Sprintf(GuidFieldMustNotExists, GuidFieldJsonTag))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsExceptSelfByGuid(c *gin.Context, guidColName string, guidValue interface{}, id uint) error {
	if r.CheckResourceExistsExceptSelfByGuid(guidColName, guidValue, id) {
		err := errors.New(fmt.Sprintf(ResourceMustNotExistsExceptSelf, r.Name))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsExceptSelfByModelPtrWithGuid(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string, id uint) error {
	if r.CheckResourceExistsExceptSelfByModelPtrWithGuid(modelPtr, GuidFieldJsonTag, id) {
		err := errors.New(fmt.Sprintf(ResourceMustNotExistsExceptSelf, r.Name))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}
