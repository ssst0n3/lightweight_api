package test

import (
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"gorm.io/gorm"
)

func InitEmptyUser(DB *gorm.DB) (err error) {
	err = DB.Migrator().DropTable(&model.User{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&model.User{})
	return
}
