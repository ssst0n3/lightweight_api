package lightweight_api

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func init() {
	InitLogger()
}

func InitLogger() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	Logger.SetLevel(logrus.InfoLevel)
}
