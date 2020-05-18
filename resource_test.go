package lightweight_api

import (
	"github.com/gin-gonic/gin"
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

func TestResource_ListResources(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath, challenge.ListResources)
	challenge.TestResourceListResources(t, router)
}
