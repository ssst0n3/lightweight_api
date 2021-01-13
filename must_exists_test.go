package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResource_MustResourceExistsGetModelByIdAutoParseParam(t *testing.T) {
	router := gin.Default()
	var resource test_data.Challenge
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		_, model, _ := challenge.MustResourceExistsGetModelByIdAutoParseParam(context)
		if model != nil {
			resource = model.(test_data.Challenge)
		}
	})
	challenge.TestResourceMustResourceExistsById(t, router, test_data.Challenge1.Challenge)
	assert.Equal(t, test_data.Challenge1.Challenge, resource)
}

func TestResource_MustResourceExistsById(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		//noinspection GoUnhandledErrorResult
		challenge.MustResourceExistsByIdAutoParseParam(context)
	})
	challenge.TestResourceMustResourceExistsById(t, router, test_data.Challenge1.Challenge)
}
