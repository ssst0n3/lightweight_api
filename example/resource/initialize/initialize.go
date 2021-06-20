package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/kv_config/model"
	"github.com/ssst0n3/lightweight_api/example/resource/user"
	"net/http"
	"strconv"
)

const (
	MsgTheDataShouldNotBeInitializedNow = "The data should not be initialized now."
)

var Resource = lightweight_api.Resource{
	BaseRelativePath: lightweight_api.BaseRelativePathV1("initialize"),
}

var (
	FlagUseKVConfig  bool
	ShouldInitialize bool
)

func CheckShouldInitialize() (shouldInitialize bool, err error) {
	//value, err := .KVGetValueByKey(
	//	TableNameConfig, ColumnNameConfigValue, ColumnNameConfigKey, "is_initialized",
	//)
	config := model.Config{Key: "is_initialized"}
	err = lightweight_api.DB.First(&config).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if isInitialized, err := strconv.ParseBool(config.Value); err != nil {
		awesome_error.CheckDebug(err)
		shouldInitialize = true
		err = nil
	} else {
		shouldInitialize = !isInitialized
	}
	return
}

func Should(c *gin.Context) {
	if !FlagUseKVConfig {
		ShouldInitialize = cipher.IsInitKey
	} else {
		var err error
		ShouldInitialize, err = CheckShouldInitialize()
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
		// update to true or false?
		err := lightweight_api.DB.FirstOrCreate(&model.Config{Key: "is_initialized", Value: "true"}).Error
		if err != nil {
			lightweight_api.HandleInternalServerError(c, err)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
