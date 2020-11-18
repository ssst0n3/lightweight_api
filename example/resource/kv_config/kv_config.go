package kv_config

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/response"
	"github.com/ssst0n3/lightweight_db"
	"net/http"
)

const (
	ResourceName = "config"
)

var Resource = lightweight_api.Resource{
	Name:             ResourceName,
	TableName:        ResourceName,
	BaseRelativePath: "/api/v1/" + ResourceName,
	Model:            lightweight_db.Config{},
	GuidFieldJsonTag: lightweight_db.ColumnNameConfigKey,
}

func Key(c *gin.Context) (key string, err error) {
	key = c.Param("key")
	err = Resource.MustResourceExistsByGuid(c, lightweight_db.ColumnNameConfigKey, key)
	return
}

func GetValueByKey(key string) (value string, err error) {
	value, err = lightweight_api.Conn.KVGetValueByKey(Resource.TableName, lightweight_db.ColumnNameConfigValue, lightweight_db.ColumnNameConfigKey, key)
	awesome_error.CheckErr(err)
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
	c.JSON(http.StatusOK, value)
}

func Delete(c *gin.Context) {
	key, err := Key(c)
	if err != nil {
		return
	}
	if err := lightweight_api.Conn.DeleteObjectByGuid(Resource.TableName, lightweight_db.ColumnNameConfigKey, key); err != nil {
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
	var model lightweight_db.Config
	if err := c.ShouldBindJSON(&model); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}
	if err := lightweight_api.Conn.UpdateObjectSingleColumnByGuid(
		lightweight_db.ColumnNameConfigKey,
		key,
		lightweight_db.TableNameConfig,
		lightweight_db.ColumnNameConfigValue,
		model.Value,
	); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
}
