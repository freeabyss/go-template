package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/spf13/viper"
	"github.com/Sirupsen/logrus"
	"maizuo.com/back-end/go-template/src/server/common"
	"log"
	"flag"
)

func SetUpLoggerForTest() {
	isDevelopment := viper.GetBool("isDevelopment")
	if isDevelopment {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stderr)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logFilePath := viper.GetString("log.path")
		fmt.Println("logPath:",logFilePath)
		logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		fmt.Println("logPath:err:",err)
		if err != nil {
			logrus.Fatalf("open file error :%s \n", logFilePath)
			//TeardownLogger()
		}
		logrus.SetLevel(logrus.WarnLevel)
		logrus.SetOutput(logFile)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	common.Logger = logrus.WithFields(logrus.Fields{})
}


func InitConfigForTest(path string) {
	fmt.Println("path:",path)
	//../../../../config/local
	flag.Parse()

	fmt.Println(path)
	viper.SetConfigName(path)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Read config fail:", err.Error())
	}
}


func GetRelToConfigPath(env string) string {
	dir, err := filepath.Abs(filepath.Dir("."))
	fmt.Println("dir::::=",dir)
	dirS := strings.Split(dir, "\\")
	n := 0
	for i := len(dirS) - 1; i >= 0; i-- {
		//fmt.Println(dirS[i])
		n++
		if dirS[i] == "src" {
			break
		}
	}
	path := ""
	for i := 0; i < n; i++ {
		path += "../"
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return fmt.Sprint(path+"config/") + env
}
