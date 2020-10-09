package lightweight_api

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/ssst0n3/lightweight_db"
	"github.com/ssst0n3/lightweight_db/example/sqlite"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	//_ "modernc.org/sqlite"
	"os"
	"testing"
)

var challenge = Resource{
	Name:             "challenge",
	TableName:        "challenge",
	BaseRelativePath: "/api/v1/challenge",
	Model:            test_data.Challenge{},
	GuidFieldJsonTag: test_data.ColumnNameChallengeName,
}

func init() {
	awesome_error.CheckFatal(os.Setenv(consts.EnvDirSecret, "/tmp/secret"))
	awesome_error.CheckFatal(os.Setenv(lightweight_db.EnvDbDsn, "test/test_data/base.sqlite"))
	Conn = sqlite.Conn()
}

func TestResource_ListResource(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath, challenge.ListResource)
	challenge.TestResourceListResource(t, router)
}

func TestResource_MustResourceNotExistsByGuid(t *testing.T) {
	challenge.TestResourceMustResourceNotExistsByGuid(
		t,
		test_data.Challenge1.Challenge,
		test_data.ColumnNameChallengeName,
		test_data.Challenge1.Name,
	)
}

func TestResource_CreateResourceNew(t *testing.T) {
	awesome_reflect.MustNotPointer(challenge.Model)
	modelPtr := awesome_reflect.EmptyPointerOfModel(challenge.Model)
	assert.NoError(t, json.Unmarshal([]byte(`{"name":"name"}`), modelPtr))
	spew.Dump(modelPtr)
}

func TestResource_CreateResource(t *testing.T) {
	router := gin.Default()
	router.POST(challenge.BaseRelativePath, challenge.CreateResource)
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
	router.PUT(challenge.BaseRelativePath+"/:id", challenge.UpdateResource)
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
