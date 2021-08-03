package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	_ "github.com/ssst0n3/lightweight_api/response" // for swaggo
	"net/http"
)

const (
	ResourceName = "user"
)

//var Resource = lightweight_api.Resource{
//	Name:             ResourceName,
//	TableName:        ResourceName,
//	BaseRelativePath: lightweight_api.BaseRelativePathV1(ResourceName),
//	User:            User{},
//	GuidFieldJsonTag: "username",
//}

var Resource = lightweight_api.NewResource(ResourceName, model.SchemaUser.Table, model.User{}, model.SchemaUser.FieldsByName["Username"].DBName)

// List godoc
// @Summary list user
// @Description return users
// @Tags User
// @ID list-user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} model.ListUserBody
// @Router /api/v1/user [get]
func List(c *gin.Context) {
	var users []model.ListUserBody
	err := lightweight_api.DB.Select(
		model.SchemaUser.FieldsByName["ID"].DBName,
		model.SchemaUser.FieldsByName["Username"].DBName,
		model.SchemaUser.FieldsByName["IsAdmin"].DBName,
	).Table(Resource.TableName).Find(&users).Error

	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// Create godoc
// @Summary create user
// @Description Add a user
// @Tags User
// @ID create-user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body model.CreateUserBody true "Create User"
// @Success 200 {object} response.CreateSuccess
// @Router /api/v1/user [post]
func Create(c *gin.Context) {
	Resource.CreateResourceTemplate(c, func(modelPtr interface{}) (err error) {
		u := modelPtr.(*model.User)
		//u.Password, err = cipher.CommonCipher.Encrypt([]byte(u.Password))
		err = EncryptUser(u)
		return
	}, nil)
}

// AnonymousCreate godoc
// @Summary create user without authentication if table user is empty
// @Description Add a user
// @Tags User
// @ID create-user
// @Accept json
// @Produce json
// @Param user body model.CreateUserBody true "Create User"
// @Success 200 {object} response.CreateSuccess
// @Router /api/v1/user/init [post]
func AnonymousCreate(c *gin.Context) {
	// check exists
	var count int64
	err := lightweight_api.DB.Model(&model.User{}).Count(&count).Error
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if count == 0 {
		Resource.CreateResourceTemplate(c, func(modelPtr interface{}) (err error) {
			u := modelPtr.(*model.User)
			//u.Password, err = cipher.CommonCipher.Encrypt([]byte(u.Password))
			err = EncryptUser(u)
			return
		}, nil)
	} else {
		c.Status(http.StatusForbidden)
		return
	}
}

// UpdateBasic godoc
// @Summary update basic
// @Description updates some basic information of user
// @Tags User
// @ID update-basic
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param user_update_basic_body body model.UpdateBasicBody true "Update user's basic information"
// @Success 200 {object} response.Base
// @Router /api/v1/user/{id}/basic [put]
func UpdateBasic(c *gin.Context) {
	Resource.UpdateResourceTemplate(c, model.UpdateBasicBody{}, nil)
}

// UpdatePassword godoc
// @Summary update password
// @Description update user's password
// @Tags User
// @ID update-password
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param user_update_password_body body model.UpdatePasswordBody true "Update user's password"
// @Success 200 {object} response.Base
// @Router /api/v1/user/{id}/password [put]
func UpdatePassword(c *gin.Context) {
	Resource.UpdateResourceTemplate(c, model.UpdatePasswordBody{}, func(modelPtr interface{}) (err error) {
		u := modelPtr.(*model.User)
		u.Password, err = cipher.CommonCipher.Encrypt([]byte(u.Password))
		return
	})
}
