package kv_config

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/kv_config/model"
	"github.com/ssst0n3/lightweight_api/response"
	"net/http"
)

const (
	ResourceName = "config"
)

var Resource = lightweight_api.Resource{
	Name:             ResourceName,
	TableName:        ResourceName,
	BaseRelativePath: lightweight_api.BaseRelativePathV1(ResourceName),
	Model:            model.Config{},
	GuidFieldJsonTag: model.SchemaConfig.FieldsByName["Key"].DBName,
}

func Key(c *gin.Context) (key string, err error) {
	key = c.Param("key")
	err = Resource.MustResourceExistsByGuid(c, model.SchemaConfig.FieldsByName["Key"].DBName, key)
	return
}

func GetValueByKey(key string) (value string, err error) {
	var config model.Config
	err = lightweight_api.DB.Where(&model.Config{Key: key}).First(&config).Error
	awesome_error.CheckErr(err)
	value = config.Value
	return
}

func Get(c *gin.Context) {
	key, err := Key(c)
	if err != nil {
		return
	}

	value, err := GetValueByKey(key)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.String(http.StatusOK, value)
}

func Delete(c *gin.Context) {
	key, err := Key(c)
	if err != nil {
		return
	}

	err = lightweight_api.DB.Unscoped().Delete(&model.Config{Key: key}).Error
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	response.DeleteSuccess200(c)
}

func Update(c *gin.Context) {
	key, err := Key(c)
	if err != nil {
		return
	}
	var m model.Config
	if err := c.ShouldBindJSON(&m); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	err = lightweight_api.DB.Model(&model.Config{Key: key}).Updates(&model.Config{Key: key, Value: m.Value}).Error
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
}
