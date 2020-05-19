package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	_ "modernc.org/sqlite"
	"testing"
)

var challenge = Resource{
	Name:             "challenge",
	TableName:        "challenge",
	BaseRelativePath: "/api/v1/challenge",
}

func init() {
	InitConnector("sqlite", "test/test_data/base.sqlite")
}

func TestResource_ListResource(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath, challenge.ListResource)
	challenge.TestResourceListResource(t, router)
}

func TestResource_CreateResource(t *testing.T) {
	router := gin.Default()
	router.POST(challenge.BaseRelativePath, func(context *gin.Context) {
		challenge.CreateResource(context, &test_data.Challenge{}, test_data.ColumnNameChallengeName)
	})
	challenge.TestResourceCreateResource(
		t, router,
		test_data.Challenge1.Challenge,
		test_data.ColumnNameChallengeName,
		test_data.Challenge1.Name,
	)
}
