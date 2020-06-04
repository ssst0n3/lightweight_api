package lightweight_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func (r *Resource) TestResourceMustResourceExistsById(t *testing.T, router *gin.Engine, resource interface{}) {
	t.Run("not exists", func(t *testing.T) {
		r.DeleteAllObjects()
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
