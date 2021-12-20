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
	if level, err := logrus.ParseLevel(c.Log.Level); err == nil {
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
		}
		Logger.SetOutput(src)

		logFile, err := os.OpenFile(fileName + ".error", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
		if err != nil {
			fmt.Println("init error log err:", err)
			os.Exit(1)
		}
		Logger.AddHook(&StouchHook{LogFile: logFile})
	}
}


type StouchHook struct {
	LogFile *os.File
}

// Levels 只定义 error 和 panic 等级的日志,其他日志等级不会触发 hook
func (h *StouchHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

// Fire 将异常日志写入到指定日志文件中
func (h *StouchHook) Fire(entry *logrus.Entry) error {
	message, _ := entry.String()
	if _, err := h.LogFile.Write([]byte(message)); err != nil {
		return err
	}
	return nil
}
