package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user/test"
	"github.com/ssst0n3/lightweight_api/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestGetUserByJwt(t *testing.T) {
	middleware.JwtSecret = []byte("example")
	assert.NoError(t, test.InitEmptyUser(lightweight_api.DB))
	assert.NoError(t, lightweight_api.DB.Create(&test.UserAdmin).Error)
	token, err := middleware.GenerateToken(test.UserAdmin.ID, true, 3*time.Hour)
	c := &gin.Context{
		Request: &http.Request{
			Header: http.Header{
				"Cookie": []string{"token=" + token},
			},
		},
	}
	assert.NoError(t, err)
	user, err := GetUserByJwt(c)
	assert.NoError(t, err)
	assert.Equal(t, test.UserAdmin, user)
}
