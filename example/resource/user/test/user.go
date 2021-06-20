package test

import (
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"gorm.io/gorm"
)

var UserAdmin = model.User{
	Model: gorm.Model{
		ID: 1,
	},
	UpdateBasicBody: model.UpdateBasicBody{
		Username: "admin",
		IsAdmin:  true,
	},
	UpdatePasswordBody: model.UpdatePasswordBody{
		Password: "admin",
	},
}
