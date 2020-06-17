package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetToken(t *testing.T) {
	token := "token"
	t.Run("cookie", func(t *testing.T) {
		c := &gin.Context{
			Request: &http.Request{
				Header: http.Header{
					"Cookie": []string{"token=" + token},
				},
			},
		}
		got, err := GetToken(c)
		assert.NoError(t, err)
		assert.Equal(t, token, got)
	})
	t.Run("header", func(t *testing.T) {
		c := &gin.Context{
			Request: &http.Request{
				Header: http.Header{},
			},
		}
		c.Request.Header.Set("token", token)
		got, err := GetToken(c)
		assert.NoError(t, err)
		assert.Equal(t, token, got)
	})
	t.Run("nothing", func(t *testing.T) {
		c := &gin.Context{
			Request: &http.Request{
			},
		}
		_, err := GetToken(c)
		assert.Error(t, err)
	})
	t.Run("both", func(t *testing.T) {
		c := &gin.Context{
			Request: &http.Request{
				Header: http.Header{
					"Cookie": []string{"token=" + token},
				},
			},
		}
		c.Request.Header.Set("token", token)
		got, err := GetToken(c)
		assert.NoError(t, err)
		assert.Equal(t, token, got)
	})
}
