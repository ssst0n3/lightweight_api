package middleware

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/awesome_libs/secret"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"os"
)

const (
	FilenameJwtKey = "lightweight_api"
	HintInitData   = "Hint: you need init your data, because the JwtSecret is init key."
)

func InitJwtKey() {
	if len(os.Getenv(consts.EnvDirSecret)) > 0 {
		var err error
		JwtSecret, IsInitKey, err = secret.LoadKey(FilenameJwtKey)
		awesome_error.CheckFatal(err)
		log.Logger.Debug(HintInitData)
	}
}

func init() {
	InitJwtKey()
}
