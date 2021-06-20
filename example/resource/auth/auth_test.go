package auth

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"github.com/ssst0n3/lightweight_api/example/resource/user/test"
	"github.com/ssst0n3/lightweight_api/middleware"
	"github.com/ssst0n3/lightweight_api/test/test_config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	test_config.Init()
	awesome_error.CheckFatal(lightweight_api.InitGormDB())
	cipher.Init()
	middleware.InitJwtKey()
}

func TestLogin(t *testing.T) {
	assert.NoError(t, test.InitEmptyUser(lightweight_api.DB))
	passwordPlain := test.UserAdmin.Password
	assert.NoError(t, user.EncryptUser(&test.UserAdmin))
	lightweight_api.DB.Create(&test.UserAdmin)

	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		u := model.User{
			UpdateBasicBody: model.UpdateBasicBody{
				Username: test.UserAdmin.Username,
			},
			UpdatePasswordBody: model.UpdatePasswordBody{
				Password: passwordPlain,
			},
		}
		marshal, err := json.Marshal(u)
		assert.NoError(t, err)
		c.Request, err = http.NewRequest(http.MethodGet, "", bytes.NewReader(marshal))
		Login(c)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("wrong password", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		u := model.User{
			UpdateBasicBody: model.UpdateBasicBody{
				Username: test.UserAdmin.Username,
			},
			UpdatePasswordBody: model.UpdatePasswordBody{
				Password: "wrong",
			},
		}
		marshal, err := json.Marshal(u)
		assert.NoError(t, err)
		c.Request, err = http.NewRequest(http.MethodGet, "", bytes.NewReader(marshal))
		Login(c)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), wrongUsernameOrPassword)
	})
}
