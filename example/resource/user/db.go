package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
	"github.com/ssst0n3/lightweight_api/middleware"
)

func EncryptUser(user *model.User) (err error) {
	user.Password, err = cipher.CommonCipher.Encrypt([]byte(user.Password))
	return
}

func GetUserByJwt(c *gin.Context) (user model.User, err error) {
	token, err := middleware.GetToken(c)
	if err != nil {
		return
	}
	claims, err := middleware.ParseToken(token)
	if err != nil {
		return
	}
	userId := claims.UserId
	err = lightweight_api.DB.Model(&model.User{}).First(&user, userId).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
