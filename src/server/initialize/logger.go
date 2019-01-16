package initialize

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"

	"maizuo.com/back-end/go-template/src/server/common"
)

func SetupLogger() {
	isDevelopment := viper.GetBool("isDevelopment")
	if isDevelopment {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stderr)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logFilePath := viper.GetString("log.path")
		logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			logrus.Fatalf("open file error :%s \n", logFilePath)
			TeardownLogger()
		}
		logrus.SetLevel(logrus.WarnLevel)
		logrus.SetOutput(logFile)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	common.Logger = logrus.WithFields(logrus.Fields{
		"system": viper.GetString("name"),
	})
}

func TeardownLogger() {
	var LogFile *os.File
	logrus.Printf("logger file stream closed.")
	LogFile.Close()
}
