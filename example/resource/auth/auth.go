package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/middleware"
	"net/http"
	"time"
)

var AuthResource = lightweight_api.Resource{
	BaseRelativePath: "/api/v1/auth",
}

func Login(c *gin.Context) {
	var user User
	errWrong := errors.Errorf("username or password wrong")

	if err := c.BindJSON(&user); err != nil {
		lightweight_api.HandleStatusBadRequestError(c, err)
		return
	}

	var userWithId UserWithId
	err := lightweight_api.Conn.OrmShowObjectByGuidUsingReflectBind(TableNameUser, ColumnNameUserUsername, user.Username, &userWithId)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if userWithId.Id == 0 {
		lightweight_api.HandleInternalServerError(c, errWrong)
		return
	}

	check, err := middleware.CheckPassword(user.Password, userWithId.Password)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	if !check {
		lightweight_api.HandleInternalServerError(c, errWrong)
		return
	}

	token, err := middleware.GenerateToken(userWithId.Id, userWithId.IsAdmin, 3*time.Hour)
	if err != nil {
		lightweight_api.HandleInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":    token,
		"username": userWithId.Username,
		"is_admin": userWithId.IsAdmin,
		"user_id":  userWithId.Id,
	})
}
