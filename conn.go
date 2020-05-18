package lightweight_api

import (
	"github.com/ssst0n3/lightweight_db"
)

var Conn lightweight_db.Connector

func InitConnector(driverName string, dsn string) {
	Conn = lightweight_db.Connector{
		DriverName: driverName,
		Dsn:        dsn,
	}
	Conn.Init()
}
