package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"github.com/ssst0n3/lightweight_api/middleware"
	"net/http"
)

var Resource = lightweight_api.Resource{
	BaseRelativePath: lightweight_api.BaseRelativePathV1("auth"),
}

func Login(c *gin.Context) {
	var u model.User
	errWrong := errors.Errorf(wrongUsernameOrPassword)

	if err := c.BindJSON(&u); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}

	var user model.User
	err := lightweight_api.DB.Where(&model.User{
		CreateUserBody: model.CreateUserBody{
			UpdateBasicBody: model.UpdateBasicBody{
				Username: u.Username,
			},
		},
	}).Find(&user).Error

	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if user.ID == 0 {
		lightweight_api.HandleInternalServerError(c, errWrong)
		return
	}

	check, err := middleware.CheckPassword(u.Password, user.Password)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if !check {
		lightweight_api.HandleInternalServerError(c, errWrong)
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.IsAdmin, DurationToken)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":    token,
		"username": user.Username,
		"is_admin": user.IsAdmin,
		"user_id":  user.ID,
	})
}
