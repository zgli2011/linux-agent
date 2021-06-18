package log

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
	// 此处初始化配置文件
	err := config.LoadConfiguration()
	if err != nil {
		log.Fatal("load config failed")
	}
	// 获取日志相关的配置信息
	log_config := config.GetConfiguration().System.Log
	// 开启日志切割
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
	// 设置日志级别
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

// init函数先于main函数自动执行
func init() {
	Log = NewLogger()
}
