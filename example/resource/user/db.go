package user

import (
	"github.com/ssst0n3/awesome_libs/cipher"
	"github.com/ssst0n3/lightweight_api/example/resource/user/model"
)

func EncryptUser(user *model.User) (err error) {
	user.Password, err = cipher.CommonCipher.Encrypt([]byte(user.Password))
	return
}
