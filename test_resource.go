package lightweight_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func (r *Resource) DeleteAllObjects() {
	Conn.DeleteAllObjects(r.TableName)
}

func ObjectOperate(req *http.Request, router *gin.Engine) *httptest.ResponseRecorder {
	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	Logger.Infof("body: %+v", w.Body.String())
	return w
}

func (r *Resource) DeleteAllThenObjectOperate(req *http.Request, router *gin.Engine) *httptest.ResponseRecorder {
	r.DeleteAllObjects()
	w := ObjectOperate(req, router)
	return w
}

func (r *Resource) TestResourceListResource(t *testing.T, router *gin.Engine) {
	req, _ := http.NewRequest(http.MethodGet, r.BaseRelativePath, nil)
	w := ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}

func (r *Resource) TestResourceMustResourceNotExistsByGuid(t *testing.T, resource interface{}, guidColName string, guidValue interface{}) {
	t.Run("not exists", func(t *testing.T) {
		r.DeleteAllObjects()
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		err := r.MustResourceNotExistsByGuid(c, guidColName, guidValue)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("exists", func(t *testing.T) {
		r.DeleteAllObjects()
		_, err := Conn.CreateObject(r.TableName, resource)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		err = r.MustResourceNotExistsByGuid(c, guidColName, guidValue)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		expect, err := json.Marshal(gin.H{
			"success": false,
			"reason":  fmt.Sprintf(ResourceAlreadyExists, r.Name, guidColName, guidValue),
		})
		assert.NoError(t, err)
		assert.Equal(t, string(expect), w.Body.String())
	})
}

func (r *Resource) TestResourceCreateResource(t *testing.T, router *gin.Engine, obj interface{}, guidColName string, guidValue interface{}) {
	objJson, err := json.Marshal(obj)
	assert.Equal(t, nil, err)
	t.Run("not exists", func(t *testing.T) {
		reader := strings.NewReader(string(objJson))
		req, _ := http.NewRequest(http.MethodPost, r.BaseRelativePath, reader)
		w := r.DeleteAllThenObjectOperate(req, router)
		assert.Equal(t, http.StatusOK, w.Code)
		exists, err := Conn.IsResourceExistsByGuid(r.TableName, guidColName, guidValue)
		assert.NoError(t, err)
		assert.Equal(t, true, exists)
	})
	t.Run("exists", func(t *testing.T) {
		reader := strings.NewReader(string(objJson))
		req, _ := http.NewRequest(http.MethodPost, r.BaseRelativePath, reader)
		w := ObjectOperate(req, router)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectBytes, err := json.Marshal(gin.H{
			"success": false,
			"reason":  fmt.Sprintf(GuidFieldMustNotExists, guidColName),
		})
		assert.NoError(t, err)
		assert.Contains(t, string(expectBytes), w.Body.String())
	})
}

func (r *Resource) TestResourceDeleteResource(t *testing.T, router *gin.Engine, obj interface{}) {
	r.DeleteAllObjects()
	id, err := Conn.CreateObject(r.TableName, obj)
	assert.NoError(t, err)
	req, _ := http.NewRequest(http.MethodDelete, r.BaseRelativePath+"/"+strconv.FormatInt(id, 10), nil)
	w := ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	exists := Conn.IsResourceExistsById(r.TableName, id)
	assert.Equal(t, false, exists)
}

func (r *Resource) TestResourceUpdateResource(t *testing.T, router *gin.Engine, obj interface{}, objUpdate interface{}, modelPtr interface{}) {
	id, err := Conn.CreateObject(r.TableName, obj)
	assert.NoError(t, err)
	objJson, err := json.Marshal(objUpdate)
	reader := strings.NewReader(string(objJson))
	req, _ := http.NewRequest(http.MethodPut, r.BaseRelativePath+"/"+strconv.FormatInt(id, 10), reader)
	w := ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}

func (r *Resource) TestResourceShowResource(t *testing.T, router *gin.Engine, obj interface{}) {
	id, err := Conn.CreateObject(r.TableName, obj)
	assert.NoError(t, err)
	req, _ := http.NewRequest(http.MethodGet, r.BaseRelativePath+"/"+strconv.FormatInt(id, 10), nil)
	w := ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}
