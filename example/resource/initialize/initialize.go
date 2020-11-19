package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user"
	"github.com/ssst0n3/lightweight_db"
	"net/http"
)

const (
	MsgTheDataShouldNotBeInitializedNow = "The data should not be initialized now."
)

var Resource = lightweight_api.Resource{
	BaseRelativePath: "/api/v1/initialize",
}

var (
	FlagUseKVConfig  bool
	ShouldInitialize bool
)

func Should(c *gin.Context) {
	if !FlagUseKVConfig {
		ShouldInitialize = cipher.IsInitKey
	} else {
		var err error
		ShouldInitialize, err = lightweight_api.Conn.ShouldInitialize()
		if err != nil {
			lightweight_api.HandleInternalServerError(c, err)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "should_initialize": ShouldInitialize})
}

func CreateUser(c *gin.Context) {
	if cipher.IsInitKey {
		user.Create(c)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     MsgTheDataShouldNotBeInitializedNow,
		})
	}
}

func End(c *gin.Context) {
	cipher.IsInitKey = false
	if !FlagUseKVConfig {
		ShouldInitialize = cipher.IsInitKey
	} else {
		ShouldInitialize = false
		if err := lightweight_api.Conn.UpdateObjectSingleColumnByGuid(
			lightweight_db.ColumnNameConfigKey,
			"is_initialized",
			lightweight_db.TableNameConfig,
			lightweight_db.ColumnNameConfigValue,
			"false",
		); err != nil {
			lightweight_api.HandleInternalServerError(c, err)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
