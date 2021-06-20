package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"github.com/ssst0n3/lightweight_api/example/resource/user/test"
	"github.com/ssst0n3/lightweight_api/test/test_config"
	"github.com/stretchr/testify/assert"
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
