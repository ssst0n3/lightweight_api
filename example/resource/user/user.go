package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api"
	"net/http"
)

const (
	ResourceName = "user"
)

var Resource = lightweight_api.Resource{
	Name:             ResourceName,
	TableName:        ResourceName,
	BaseRelativePath: "/api/v1/" + ResourceName,
	Model:            Model{},
	GuidFieldJsonTag: "username",
}

type UpdateBasicBody struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

/*
need encrypt
*/
type UpdatePasswordBody struct {
	Password string `json:"password"`
}

type Model struct {
	UpdateBasicBody
	UpdatePasswordBody
}

type ModelWithId struct {
	Id uint `json:"id"`
	Model
}

const (
	ColumnNameUsername = "username"
	ColumnNameIsAdmin  = "is_admin"
)

func List(c *gin.Context) {
	query := awesome_libs.Format("SELECT id, {.username}, {.is_admin}", awesome_libs.Dict{
		"username": ColumnNameUsername,
		"is_admin": ColumnNameIsAdmin,
	})
	if objects, err := lightweight_api.Conn.ListObjects(query); err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, objects)
	}
}

func Create(c *gin.Context) {
	Resource.CreateResourceTemplate(c, func(modelPtr interface{}) (err error) {
		u := modelPtr.(*Model)
		u.Password, err = cipher.CommonCipher.Encrypt([]byte(u.Password))
		return
	}, nil)
}

func UpdateBasic(c *gin.Context) {
	Resource.UpdateResourceTemplate(c, UpdateBasicBody{}, nil)
}

func UpdatePassword(c *gin.Context) {
	Resource.UpdateResourceTemplate(c, UpdatePasswordBody{}, func(modelPtr interface{}) (err error) {
		u := modelPtr.(*Model)
		u.Password, err = cipher.CommonCipher.Encrypt([]byte(u.Password))
		return
	})
}
