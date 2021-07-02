package user

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"github.com/ssst0n3/lightweight_api/example/resource/user/test"
	"github.com/ssst0n3/lightweight_api/test/test_config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	test_config.Init()
	awesome_error.CheckFatal(lightweight_api.InitGormDB())
}

func TestList(t *testing.T) {
	assert.NoError(t, test.InitEmptyUser(lightweight_api.DB))
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	lightweight_api.DB.Create(&test.UserAdmin)
	List(c)
	log.Logger.Info(w.Body.String())
	var users []model.User
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &users))
	assert.Equal(t, test.UserAdmin.ID, users[0].ID)
	assert.Equal(t, test.UserAdmin.Username, users[0].Username)
	assert.Equal(t, "", users[0].Password)
}

func TestAnonymousCreate(t *testing.T) {
	cipher.Init()
	router := gin.Default()
	router.POST("/create", AnonymousCreate)
	assert.NoError(t, test.InitEmptyUser(lightweight_api.DB))
	body, err := json.Marshal(test.UserAdmin)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
	w := lightweight_api.ObjectOperate(req, router)
	assert.Equal(t, http.StatusOK, w.Code)
	var user model.User
	assert.NoError(t, lightweight_api.DB.Model(&user).First(&user).Error)
	assert.Equal(t, test.UserAdmin.Username, user.Username)
	{
		// not empty
		req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewReader(body))
		assert.NoError(t, err)
		w := lightweight_api.ObjectOperate(req, router)
		assert.Equal(t, http.StatusForbidden, w.Code)
		var count int64
		assert.NoError(t, lightweight_api.DB.Table(model.SchemaUser.Table).Count(&count).Error)
		assert.Equal(t, int64(1), count)
	}
}
