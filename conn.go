package lightweight_api

import (
	"bytes"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/sqlite"
	//"github.com/cloudquery/sqlite"
	"github.com/glebarez/sqlite"
	"os"
)

const (
	EnvDriverName     = "DB_DRIVER_NAME"
	EnvDbDsn          = "DB_DSN"
	EnvDbName         = "DB_NAME"
	EnvDbHost         = "DB_HOST"
	EnvDbPort         = "DB_PORT"
	EnvDbUser         = "DB_USER"
	EnvDbPasswordFile = "DB_PASSWORD_FILE"
)

var DB *gorm.DB

func GetDsnFromEnvNormal() (dsn string) {
	if dsn = os.Getenv(EnvDbDsn); len(dsn) == 0 {
		dbProtocol := "tcp"
		dbName := os.Getenv(EnvDbName)
		dbHost := os.Getenv(EnvDbHost)
		dbPort := os.Getenv(EnvDbPort)
		dbUser := os.Getenv(EnvDbUser)
		dbPasswordFile := os.Getenv(EnvDbPasswordFile)
		password, err := os.ReadFile(dbPasswordFile)
		if err != nil {
			panic(err)
		}
		password = bytes.TrimSpace(password)

		dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?collation=utf8mb4_general_ci&maxAllowedPacket=0&parseTime=true", dbUser, password, dbProtocol, dbHost, dbPort, dbName)
	}
	return
}

func InitGormDB() (err error) {
	driverName := os.Getenv(EnvDriverName)
	dsn := GetDsnFromEnvNormal()
	switch driverName {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	default:
		log.Logger.Warnf(HintDriverNameNotRecognized, driverName)
	}
	return
}
