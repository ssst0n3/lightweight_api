package lightweight_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api/test/model"
	"github.com/ssst0n3/lightweight_api/test/test_data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func WrapperTestResourceMustResourceExistsById(t *testing.T, router *gin.Engine) {
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		req, _ := http.NewRequest(http.MethodGet, challenge.BaseRelativePath+"/1", nil)
		w := ObjectOperate(req, router)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		req, _ := http.NewRequest(http.MethodGet, challenge.BaseRelativePath+fmt.Sprintf("/%d", test_data.Challenge1.ID), nil)
		w := ObjectOperate(req, router)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestResource_MustResourceExistsGetModelByIdAutoParseParam(t *testing.T) {
	router := gin.Default()
	var resource model.Challenge
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		_, m, _ := challenge.MustResourceExistsGetModelByIdAutoParseParam(context)
		if m != nil {
			resource = m.(model.Challenge)
		}
	})
	WrapperTestResourceMustResourceExistsById(t, router)
	assert.Equal(t, test_data.Challenge1.Name, resource.Name)
}

func TestResource_MustResourceExistsByIdAutoParseParam(t *testing.T) {
	router := gin.Default()
	var id int64
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		id, _ = challenge.MustResourceExistsByIdAutoParseParam(context)
	})
	WrapperTestResourceMustResourceExistsById(t, router)
	assert.Equal(t, test_data.Challenge1.ID, uint(id))
}

func TestResource_MustResourceExistsById(t *testing.T) {
	router := gin.Default()
	router.GET(challenge.BaseRelativePath+"/:id", func(context *gin.Context) {
		_ = challenge.MustResourceExistsById(context, test_data.Challenge1.ID)
	})
	WrapperTestResourceMustResourceExistsById(t, router)
}

func WrapperTestResourceMustResourceExistsByGuid(t *testing.T, executor func(c *gin.Context) error) {
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		assert.Error(t, executor(c))
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		assert.NoError(t, executor(c))
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestResource_MustResourceExistsByGuid(t *testing.T) {
	WrapperTestResourceMustResourceExistsByGuid(t, func(c *gin.Context) error {
		return challenge.MustResourceExistsByGuid(c, challenge.GuidFieldJsonTag, test_data.Challenge1.Name)
	})
}

func WrapperTestResourceMustResourceNotExistsByGuid(t *testing.T, executor func(c *gin.Context) error) {
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		assert.NoError(t, executor(c))
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		assert.Error(t, executor(c))
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestResource_MustResourceNotExistsByGuid(t *testing.T) {
	WrapperTestResourceMustResourceNotExistsByGuid(t, func(c *gin.Context) error {
		return challenge.MustResourceNotExistsByGuid(c, challenge.GuidFieldJsonTag, test_data.Challenge1.Name)
	})
}

func TestResource_MustResourceNotExistsByModelPtrWithGuid(t *testing.T) {
	WrapperTestResourceMustResourceNotExistsByGuid(t, func(c *gin.Context) error {
		return challenge.MustResourceNotExistsByModelPtrWithGuid(c, &test_data.Challenge1, challenge.GuidFieldJsonTag)
	})
}

func WrapperTestResourceMustResourceNotExistsExceptSelf(t *testing.T, executor func(c *gin.Context) error) {
	t.Run("not exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		assert.NoError(t, executor(c))
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("not exists2", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		assert.NoError(t, executor(c))
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("exists", func(t *testing.T) {
		assert.NoError(t, test_data.InitEmptyChallenge(DB))
		DB.Create(&test_data.Challenge1)
		DB.Create(&test_data.ChallengeSameName)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		assert.Error(t, executor(c))
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestResource_MustResourceNotExistsExceptSelfByGuid(t *testing.T) {
	WrapperTestResourceMustResourceNotExistsExceptSelf(t, func(c *gin.Context) error {
		return challenge.MustResourceNotExistsExceptSelfByGuid(c, challenge.GuidFieldJsonTag, test_data.Challenge1.Name, test_data.Challenge1.ID)
	})
}

func TestResource_MustResourceNotExistsExceptSelfByModelPtrWithGuid(t *testing.T) {
	WrapperTestResourceMustResourceNotExistsExceptSelf(t, func(c *gin.Context) error {
		return challenge.MustResourceNotExistsExceptSelfByModelPtrWithGuid(c, &test_data.Challenge1, challenge.GuidFieldJsonTag, test_data.Challenge1.ID)
	})
}