package lightweight_api

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_db"
	"github.com/ssst0n3/lightweight_db/example/mysql"
	"github.com/ssst0n3/lightweight_db/example/sqlite"
	"os"
)

var Conn lightweight_db.Connector

func InitConnector() {
	driverName := os.Getenv(EnvDriverName)
	switch driverName {
	case "mysql":
		Conn = mysql.Conn()
	case "sqlite":
		Conn = sqlite.Conn()
	default:
		log.Logger.Debugf(HintDriverNameNotRecognized, driverName)
	}
}

func init() {
	InitConnector()
}
