package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user"
	"net/http"
)

const (
	MsgTheDataShouldNotBeInitializedNow = "The data should not be initialized now."
)

var Resource = lightweight_api.Resource{
	BaseRelativePath: "/api/v1/initialize",
}

<<<<<<< HEAD
func ShouldInitialize(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "should_initialize": cipher.IsInitKey})
=======
var (
	FlagUseShouldInitialize bool
	ShouldInitialize        bool
)

func IsInitialize(c *gin.Context) {
	if !FlagUseShouldInitialize {
		ShouldInitialize = cipher.IsInitKey
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "initialize": ShouldInitialize})
>>>>>>> 30daa8c4e031040d18737e79720550ea0b46e0da
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
	if !FlagUseShouldInitialize {
		ShouldInitialize = cipher.IsInitKey
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
