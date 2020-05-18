package lightweight_api

import (
	"github.com/ssst0n3/lightweight_db"
)

var Conn lightweight_db.Connector

func InitExample(dsn string) {
	Conn = lightweight_db.Connector{
		DriverName: "mysql",
		Dsn:        dsn,
	}
	Conn.Init()
}
