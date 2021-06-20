package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	model2 "github.com/ssst0n3/lightweight_api/example/resource/kv_config/model"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func InitEmptyInitialize(DB *gorm.DB) (err error) {
	err = DB.Migrator().DropTable(&model.User{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&model.User{})
	return
}

func TestEnd(t *testing.T) {
	t.Run("not use kv", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		End(c)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("use kv", func(t *testing.T) {
		FlagUseKVConfig = true
		assert.NoError(t, lightweight_api.DB.AutoMigrate(&model2.Config{}))
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		End(c)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
