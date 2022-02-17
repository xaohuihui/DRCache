package common

import (
	"fmt"
	goFileRotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"strings"
	"time"
)

// author: xaohuihui
// datetime: 2022/2/17 14:43:43
// software: GoLand

// CustomizeFormat 自定义格式
type CustomizeFormat struct{}

func (c CustomizeFormat) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] %s \n",
		time.Now().Local().Format("2006-01-02 15:04:05"),
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)
	return []byte(msg), nil
}

func Log2FileByClass() {
	lfhook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: splitCofig("debug"),
		logrus.InfoLevel:  splitCofig("info"),
		logrus.WarnLevel:  splitCofig("warn"),
		logrus.ErrorLevel: splitCofig("error"),
		logrus.FatalLevel: splitCofig("fatal"),
		logrus.PanicLevel: splitCofig("panic"),
	}, LoggerClient.Formatter)
	LoggerClient.AddHook(lfhook)
}

func splitCofig(level string) *goFileRotatelogs.RotateLogs {
	// 拼接日志名
	logFile := path.Join(LogrusConfigInstance.Path, level)
	logs, err := goFileRotatelogs.New(
		logFile+"-"+LogrusConfigInstance.Suffix,
		// 生成软链， 指向最新日志文件
		goFileRotatelogs.WithLinkName(logFile),
	)
	BusErrorInstance.ThrowError(err)
	return logs
}
