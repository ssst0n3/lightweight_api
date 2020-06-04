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
	Model:            test_data.ChallengeWithId{},
}

func init() {
	InitConnector("sqlite", "test/test_data/base.sqlite")
}

func TestResource_ListResource(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath, challenge.ListResource)
	challenge.TestResourceListResource(t, router)
}

func TestResource_CheckResourceExistsById(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		//noinspection GoUnhandledErrorResult
		challenge.MustResourceExistsById(context)
	})
	challenge.TestResourceCheckResourceExistsById(t, router, test_data.Challenge1.Challenge)
}

func TestResource_CheckResourceExistsByGuid(t *testing.T) {
	challenge.TestResourceCheckResourceExistsByGuid(
		t,
		test_data.Challenge1.Challenge,
		test_data.ColumnNameChallengeName,
		test_data.Challenge1.Name,
	)
}

func TestResource_CreateResource(t *testing.T) {
	router := gin.Default()
	router.POST(challenge.BaseRelativePath, func(context *gin.Context) {
		challenge.CreateResource(context, &test_data.Challenge{}, test_data.ColumnNameChallengeName, nil)
	})
	challenge.TestResourceCreateResource(
		t, router,
		test_data.Challenge1.Challenge,
		test_data.ColumnNameChallengeName,
		test_data.Challenge1.Name,
	)
}

func TestResource_DeleteResource(t *testing.T) {
	router := gin.Default()
	router.DELETE(challenge.BaseRelativePath+"/:id", challenge.DeleteResource)
	challenge.TestResourceDeleteResource(t, router, test_data.Challenge{})
}

// Please Delete and Reset Table by your self
func TestResource_UpdateResource(t *testing.T) {
	router := gin.Default()
	router.PUT(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		challenge.UpdateResource(context, &test_data.Challenge{}, test_data.ColumnNameChallengeName, nil)
	})
	Conn.DeleteAllObjects(test_data.TableNameChallenge)
	Conn.ResetAutoIncrementSqlite(test_data.TableNameChallenge)
	challenge.TestResourceUpdateResource(
		t, router,
		test_data.Challenge1.Challenge,
		test_data.Challenge1Update.Challenge,
		&test_data.ChallengeWithId{},
	)
}

func TestResource_ShowResource(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		challenge.ShowResource(context)
	})
	Conn.DeleteAllObjects(test_data.TableNameChallenge)
	Conn.ResetAutoIncrementSqlite(test_data.TableNameChallenge)
	challenge.TestResourceShowResource(t, router, test_data.Challenge1)
}
