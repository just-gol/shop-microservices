package utils

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.SetOutput(os.Stdout)
	Logger.SetReportCaller(true)
	//Logger.SetFormatter(&logrus.TextFormatter{
	//	FullTimestamp:   true,
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	Logger.SetFormatter(&logrus.JSONFormatter{ // json格式
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true, // 开着会更好看,控制台打印json格式，但不适合高并发场景
	})

	level := logrus.InfoLevel
	if lvl := strings.TrimSpace(os.Getenv("LOG_LEVEL")); lvl != "" {
		if parsed, err := logrus.ParseLevel(strings.ToLower(lvl)); err == nil {
			level = parsed
		}
	}
	Logger.SetLevel(level)
}
