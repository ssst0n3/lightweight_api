package middleware

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseToken(t *testing.T) {
	JwtSecret = []byte("secret")
	type MyCustomClaims struct {
		IsAdmin bool `json:"is_admin"`
		jwt.MapClaims
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		true,
		jwt.MapClaims{
			"aud": []string{""},
		},
	})
	token, err := tokenClaims.SignedString(JwtSecret)
	assert.NoError(t, err)
	c, err := ParseToken(token)
	assert.NoError(t, err)
	spew.Dump(c)
}
