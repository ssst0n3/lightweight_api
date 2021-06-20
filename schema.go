package lightweight_api

import (
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"gorm.io/gorm/schema"
	"sync"
)

func InitSchema(s *schema.Schema, dst interface{}) (err error) {
	awesome_reflect.MustPointer(dst)
	s0, err := schema.Parse(dst, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return
	}
	*s = *s0
	return
}
