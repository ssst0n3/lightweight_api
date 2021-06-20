package lightweight_api

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_api/test/test_config"
	"os"
)

func init() {
	if len(os.Args) > 1 && os.Args[1][:5] == "-test" {
		// testing mode
		test_config.Init()
	}
	awesome_error.CheckFatal(InitGormDB())
	//Conn = sqlite.Conn()
}
