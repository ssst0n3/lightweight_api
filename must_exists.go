package lightweight_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (r Resource) MustResourceExistsGetModelByIdAutoParseParam(c *gin.Context) (id int64, model interface{}, err error) {
	id, err = r.MustResourceExistsByIdAutoParseParam(c)
	if err != nil {
		return
	}
	model, err = Conn.OrmShowObjectByIdUsingReflectRet(r.TableName, id, r.Model)
	if err != nil {
		HandleInternalServerError(c, err)
		return
	}
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

func (r *Resource) MustResourceExistsById(c *gin.Context, id int64) error {
	if exists, err := r.CheckResourceExistsById(id); err != nil {
		HandleInternalServerError(c, err)
	} else {
		if !exists {
			err := errors.New(fmt.Sprintf(ResourceMustExists, r.Name))
			HandleStatusBadRequestError(c, err)
			return err
		}
	}
	return nil
}

func (r *Resource) MustResourceExistsByGuid(c *gin.Context, guidColName string, guidValue interface{}) error {
	if exists, err := r.CheckResourceExistsByGuid(guidColName, guidValue); err != nil {
		HandleInternalServerError(c, err)
		return err
	} else if !exists {
		err := errors.New(fmt.Sprintf(ResourceMustExists, r.Name))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsByGuid(c *gin.Context, guidColName string, guidValue interface{}) error {
	if exists, err := r.CheckResourceExistsByGuid(guidColName, guidValue); err != nil {
		HandleInternalServerError(c, err)
		return err
	} else if exists {
		//err := errors.New(fmt.Sprintf(ResourceAlreadyExists, r.Name, guidColName, guidValue))
		err := errors.New(fmt.Sprintf(GuidFieldMustNotExists, guidColName))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsByModelPtrWithGuid(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string) error {
	exists, err := r.CheckResourceExistsByModelPtrWithGuid(modelPtr, GuidFieldJsonTag)
	if err != nil {
		HandleInternalServerError(c, err)
		return err
	}
	if exists {
		err := errors.New(fmt.Sprintf(GuidFieldMustNotExists, GuidFieldJsonTag))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsExceptSelfByGuid(c *gin.Context, guidColName string, guidValue interface{}, id int64) error {
	if exists, err := r.CheckResourceExistsExceptSelfByGuid(guidColName, guidValue, id); err != nil {
		HandleInternalServerError(c, err)
		return err
	} else if exists {
		err := errors.New(fmt.Sprintf(ResourceMustNotExistsExceptSelf, r.Name))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}

func (r *Resource) MustResourceNotExistsExceptSelfByModelPtrWithGuid(c *gin.Context, modelPtr interface{}, GuidFieldJsonTag string, id int64) error {
	if exists, err := r.CheckResourceExistsExceptSelfByModelPtrWithGuid(modelPtr, GuidFieldJsonTag, id); err != nil {
		HandleInternalServerError(c, err)
		return err
	} else if exists {
		err := errors.New(fmt.Sprintf(ResourceMustNotExistsExceptSelf, r.Name))
		HandleStatusBadRequestError(c, err)
		return err
	}
	return nil
}
