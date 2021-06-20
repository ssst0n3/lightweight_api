package model

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"gorm.io/gorm/schema"
)

type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var SchemaConfig schema.Schema

func init() {
	awesome_error.CheckFatal(lightweight_api.InitSchema(&SchemaConfig, &Config{}))
}
