package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user"
	"net/http"
	"os"
)

const (
	MsgTheDataShouldNotBeInitializedNow = "The data should not be initialized now."
)

var Resource = lightweight_api.Resource{
	BaseRelativePath: "/api/v1/initialize",
}

func Init() {
	lightweight_api.Logger.Info("start init.")
	cipher.InitCipher()
}

func init() {
	if len(os.Getenv(consts.EnvDirSecret)) > 0 {
		Init()
	}
}

func IsInitialize(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "initialize": cipher.IsInitKey})
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
	c.JSON(http.StatusOK, gin.H{"success": true})
}
