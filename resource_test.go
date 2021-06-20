package lightweight_api

import (
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"github.com/ssst0n3/lightweight_api/test/model"
	test_data2 "github.com/ssst0n3/lightweight_api/test/test_data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"

	"testing"
)

const ChallengeResourceName = "challenge"

var challenge = Resource{
	Name:             ChallengeResourceName,
	TableName:        model.SchemaChallenge.Table,
	BaseRelativePath: BaseRelativePathV1(ChallengeResourceName),
	Model:            model.Challenge{},
	GuidFieldJsonTag: model.SchemaChallenge.FieldsByName["Name"].DBName,
}

func TestResource_ListResource(t *testing.T) {
	assert.NoError(t, test_data2.InitEmptyChallenge(DB))
	DB.Create(&test_data2.Challenge1)
	router := gin.Default()
	router.GET(challenge.BaseRelativePath, challenge.ListResource)
	challenge.TestResourceListResource(t, router)
}

func TestResource_MapResourceById(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.NoError(t, test_data2.InitEmptyChallenge(DB))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		challenge.MapResourceById(c)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "{}", w.Body.String())
	})
	t.Run("not empty", func(t *testing.T) {
		assert.NoError(t, test_data2.InitEmptyChallenge(DB))
		DB.Create(&test_data2.Challenge1)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		challenge.MapResourceById(c)
		assert.Equal(t, http.StatusOK, w.Code)
		var resource map[int64]awesome_libs.Dict
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resource))
		assert.Equal(t, test_data2.Challenge1.Name, resource[1][model.SchemaChallenge.FieldsByName["Name"].DBName])
	})
}

func TestResource_CreateResourceTemplate(t *testing.T) {
	assert.NoError(t, test_data2.InitEmptyChallenge(DB))
	router := gin.Default()
	router.POST(challenge.BaseRelativePath, func(context *gin.Context) {
		challenge.CreateResourceTemplate(context, nil, nil)
	})
	marshal, err := json.Marshal(test_data2.Challenge1)
	assert.NoError(t, err)
	body := bytes.NewReader(marshal)
	req, err := http.NewRequest(http.MethodPost, challenge.BaseRelativePath, body)
	assert.NoError(t, err)
	w := ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, challenge.CheckResourceExistsByGuid(challenge.GuidFieldJsonTag, test_data2.Challenge1.Name))
	var c model.Challenge
	DB.Model(&model.Challenge{Name: test_data2.Challenge1.Name}).First(&c)
	var response struct {
		ID uint `json:"id"`
	}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, c.ID, response.ID)
}

func TestResource_CreateResourceNew(t *testing.T) {
	awesome_reflect.MustNotPointer(challenge.Model)
	modelPtr := awesome_reflect.EmptyPointerOfModel(challenge.Model)
	assert.NoError(t, json.Unmarshal([]byte(`{"name":"name"}`), modelPtr))
	spew.Dump(modelPtr)
}

func TestResource_CreateResource(t *testing.T) {
	assert.NoError(t, test_data2.InitEmptyChallenge(DB))
	router := gin.Default()
	router.POST(challenge.BaseRelativePath, challenge.CreateResource)
	challenge.TestResourceCreateResource(
		t, router,
		test_data2.Challenge1,
		model.SchemaChallenge.FieldsByName["Name"].DBName,
		test_data2.Challenge1.Name,
	)
}

func TestResource_DeleteResource(t *testing.T) {
	router := gin.Default()
	router.DELETE(challenge.BaseRelativePath+"/:id", challenge.DeleteResource)
	challenge.TestResourceDeleteResource(t, router, &test_data2.Challenge1)
}

// Please Delete and Reset Table by your self
func TestResource_UpdateResource(t *testing.T) {
	router := gin.Default()
	router.PUT(challenge.BaseRelativePath+"/:id", challenge.UpdateResource)
	assert.NoError(t, test_data2.InitEmptyChallenge(DB))
	challenge.TestResourceUpdateResource(
		t, router,
		&test_data2.Challenge1,
		&test_data2.Challenge1Update,
	)
	// TODO: check count
}

func TestResource_ShowResource(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		challenge.ShowResource(context)
	})
	assert.NoError(t, test_data2.InitEmptyChallenge(DB))
	challenge.TestResourceShowResource(t, router, &test_data2.Challenge1)
}
