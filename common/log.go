package common

import (
	"agent/config"
	"log"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}
	err := config.LoadConfiguration()
	if err != nil {
		log.Fatal("load config failed")
	}
	log_config := config.GetConfiguration().System.Log
	writer, err := rotatelogs.New(
		log_config.Path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(log_config.Path),                                      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(log_config.MaxAge)*time.Hour),             // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(log_config.RotationTime)*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatalf("config local file system for logger error: %v", err)
	}
	Log = logrus.New()
	log_level_map := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	Log.SetLevel(log_level_map[log_config.Level])
	Log.Hooks.Add(
		lfshook.NewHook(
			lfshook.WriterMap{
				logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
				logrus.InfoLevel:  writer,
				logrus.WarnLevel:  writer,
				logrus.ErrorLevel: writer,
				logrus.FatalLevel: writer,
				logrus.PanicLevel: writer,
			},
			&logrus.TextFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			},
		))
	return Log
}

func init() {
	Log = NewLogger()
}
