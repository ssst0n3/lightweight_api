package lightweight_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/lightweight_api/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAuth_GetUserId(t *testing.T) {
	userId := uint(1)
	middleware.JwtSecret = []byte("example")
	token, err := middleware.GenerateToken(userId, true)
	c := &gin.Context{
		Request: &http.Request{
			Header: http.Header{
				"Cookie": []string{"token=" + token},
			},
		},
	}
	userIdGot, err := GetUserId(c)
	assert.NoError(t, err)
	assert.Equal(t, userId, userIdGot)
}
