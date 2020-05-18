package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ObjectOperate(req *http.Request, router *gin.Engine) *httptest.ResponseRecorder {
	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Printf("body: %+v", w.Body.String())
	return w
}

func (r *Resource) TestResourceListResources(t *testing.T, router *gin.Engine) {
	req, _ := http.NewRequest(http.MethodGet, r.BaseRelativePath, nil)
	w := ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
}
