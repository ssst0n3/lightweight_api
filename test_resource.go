package lightweight_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

func (r *Resource) TestResourceCheckResourceExistsById(t *testing.T, router *gin.Engine, resource interface{}) {
	t.Run("not exists", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, r.BaseRelativePath+"/1", nil)
		w := ObjectOperate(req, router)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("exists", func(t *testing.T) {
		r.DeleteAllObjects()
		id, err := Conn.CreateObject(r.TableName, resource)
		assert.NoError(t, err)
		req, _ := http.NewRequest(http.MethodGet, r.BaseRelativePath+fmt.Sprintf("/%d", id), nil)
		w := ObjectOperate(req, router)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func (r *Resource) TestResourceCreateResource(t *testing.T, router *gin.Engine, obj interface{}, guidColName string, guidValue interface{}) {
	objJson, err := json.Marshal(obj)
	assert.Equal(t, nil, err)
	reader := strings.NewReader(string(objJson))
	req, _ := http.NewRequest(http.MethodPost, r.BaseRelativePath, reader)
	w := r.DeleteAllThenObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	exists, err := Conn.IsResourceExistsByGuid(r.TableName, guidColName, guidValue)
	assert.NoError(t, err)
	assert.Equal(t, true, exists)
}
