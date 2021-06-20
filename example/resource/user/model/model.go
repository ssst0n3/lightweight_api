package model

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type UpdateBasicBody struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

/*
UpdatePasswordBody need encrypt
*/
type UpdatePasswordBody struct {
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	UpdateBasicBody
	UpdatePasswordBody
}

const (
	ColumnNameUsername = "username"
	ColumnNameIsAdmin  = "is_admin"
)

var SchemaUser schema.Schema

func init() {
	awesome_error.CheckFatal(lightweight_api.InitSchema(&SchemaUser, &User{}))
}

// ListUserBody ======response model========
type ListUserBody struct {
	Id uint `json:"id"`
	UpdateBasicBody
}
