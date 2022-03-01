package lightweight_api

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
)

func migrate(m interface{}) (err error) {
	awesome_reflect.MustPointer(m)
	err = DB.AutoMigrate(m)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func MigrateTables(models ...interface{}) (err error) {
	for _, m := range models {
		err = migrate(m)
		if err != nil {
			return err
		}
	}
	return
}
