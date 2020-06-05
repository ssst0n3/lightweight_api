package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"testing"
)

func TestResource_MustResourceExistsById(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		//noinspection GoUnhandledErrorResult
		challenge.MustResourceExistsById(context)
	})
	challenge.TestResourceMustResourceExistsById(t, router, test_data.Challenge1.Challenge)
}