package lightweight_api

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/lightweight_db"
	"github.com/ssst0n3/lightweight_db/example/mysql"
	"os"
)

var Conn lightweight_db.Connector

func init() {
	driverName := os.Getenv(EnvDriverName)
	switch driverName {
	case "mysql":
		Conn = mysql.Conn()
	case "sqlite":
		Conn = lightweight_db.GetNewConnector("sqlite", lightweight_db.GetDsnFromEnvNormal())
	default:
		log.Logger.Debugf(HintDriverNameNotRecognized, driverName)
	}
}
