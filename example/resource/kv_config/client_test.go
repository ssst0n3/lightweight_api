package kv_config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Get(t *testing.T) {
	c := Client{
		Address: "http://127.0.0.1:13100",
	}
	res, err := c.Get(Resource.BaseRelativePath, "is_initialized")
	assert.NoError(t, err)
	spew.Dump(res)
}
