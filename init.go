package lightweight_api

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api/test/test_config"
)

func Init() {
	if !test_config.IsTestLightweightApi() {
		awesome_error.CheckFatal(InitGormDB())
	}
}
