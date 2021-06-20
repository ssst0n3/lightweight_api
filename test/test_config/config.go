package test_config

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"os"
)

const (
	EnvDriverName     = "DB_DRIVER_NAME"
	EnvDbDsn          = "DB_DSN"
)

func Init() {
	log.Logger.Info("test_config.init")
	awesome_error.CheckFatal(os.Setenv(consts.EnvDirSecret, "/tmp/secret"))
	awesome_error.CheckFatal(os.Setenv(EnvDbDsn, "/tmp/base.sqlite"))
	awesome_error.CheckFatal(os.Setenv(EnvDriverName, "sqlite"))
}
