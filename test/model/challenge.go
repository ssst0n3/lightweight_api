package model

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
)

type Challenge struct {
	gorm.Model
	Name   string `json:"name"`
	Score  int    `json:"score"`
	Solved bool   `json:"solved"`
}

var SchemaChallenge schema.Schema

func InitSchema(s *schema.Schema, dst interface{}) (err error) {
	awesome_reflect.MustPointer(dst)
	s0, err := schema.Parse(dst, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return
	}
	*s = *s0
	return
}

func init() {
	awesome_error.CheckFatal(InitSchema(&SchemaChallenge, &Challenge{}))
}
