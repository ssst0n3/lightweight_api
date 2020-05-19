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
	_, err := Conn.Exec(fmt.Sprintf("DELETE FROM %s", r.TableName))
	if err != nil {
		Logger.Fatal(err)
	}
}

func ObjectOperate(req *http.Request, router *gin.Engine) *httptest.ResponseRecorder {
	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	Logger.Infof("body: %+v", w.Body.String())
	return w
}

func (r *Resource) DeleteAllThenObjectOperate(t *testing.T, req *http.Request, router *gin.Engine) *httptest.ResponseRecorder {
	r.DeleteAllObjects()
	w := ObjectOperate(req, router)
	return w
}

func (r *Resource) TestResourceListResource(t *testing.T, router *gin.Engine) {
	req, _ := http.NewRequest(http.MethodGet, r.BaseRelativePath, nil)
	w := ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}

func (r *Resource) TestResourceCreateResource(t *testing.T, router *gin.Engine, obj interface{}, guidColName, guidValue string) {
	objJson, err := json.Marshal(obj)
	assert.Equal(t, nil, err)
	reader := strings.NewReader(string(objJson))
	req, _ := http.NewRequest(http.MethodPost, r.BaseRelativePath, reader)
	w := r.DeleteAllThenObjectOperate(t, req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	exists, err := Conn.IsResourceExistsByGuid(r.TableName, guidColName, guidValue)
	assert.NoError(t, err)
	assert.Equal(t, true, exists)
}


