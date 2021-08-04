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

// Nothing godoc
// @Summary Nothing
// @Description Just for check permission, only user with admin permission will get 200, otherwise will get 401; 只有管理员可以获得200,普通用户会401.
// @Tags Auth
// @ID auth-nothing
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200
// @Failure 401
// @Router /api/v1/auth [get]
func Nothing(c *gin.Context) {}

// Login godoc
// @Summary Login
// @Description
// @Tags Auth
// @ID auth-login
// @Accept json
// @Produce json
// @Param user body model.LoginModel true "User"
// @Success 200 {object} LoginSuccessResponse
// @Router /api/v1/auth [post]
func Login(c *gin.Context) {
	var loginModel model.LoginModel
	errWrong := errors.Errorf(wrongUsernameOrPassword)

	if err := c.BindJSON(&loginModel); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}

	var user model.User
	err := lightweight_api.DB.Where(&model.User{
		CreateUserBody: model.CreateUserBody{
			UpdateBasicBody: model.UpdateBasicBody{
				Username: loginModel.Username,
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

	check, err := middleware.CheckPassword(loginModel.Password, user.Password)
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
	c.JSON(http.StatusOK, LoginSuccessResponse{
		Token:    token,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		UserId:   user.ID,
	})
}
