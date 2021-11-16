package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

var Logger = logrus.New()

func loadLogConf(c Configuration) {
	//设置日志级别
	if level, err := logrus.ParseLevel(c.Log.Level); err != nil {
		Logger.SetLevel(level)
	}

	Logger.SetFormatter(&logrus.JSONFormatter{})
	// 日志文件
	if c.Log.LogFileName == "stdout" {
		Logger.SetOutput(os.Stdout)
	} else {
		fileName := path.Join(c.Log.LogFilePath, c.Log.LogFileName)
		// 写入文件
		src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
		if err != nil {
			fmt.Println("init log err:", err)
			os.Exit(1)
		} else {
			Logger.SetOutput(src)
		}
	}
}
