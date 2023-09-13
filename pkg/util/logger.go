package util

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

var Logger *logrus.Logger

func InitLog() {
	if Logger != nil {
		src, _ := setOutputFile()
		// 设置输出
		Logger.Out = src
		return
	}
	// 实例化
	logger := logrus.New()

	// 设置输出
	isDev := os.Getenv("ENV") == "dev"
	if isDev {
		logger.Out = os.Stdout
	} else {
		src, _ := setOutputFile()
		logger.Out = io.MultiWriter(os.Stdout, src)
	}
	// 设置日志级别
	logger.SetLevel(logrus.InfoLevel)
	// 设置日志格式
	logger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.DateTime,
		NoColors:        !isDev,
		CallerFirst:     true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullPath), line)
		},
	})
	logger.ReportCaller = true
	Logger = logger
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
