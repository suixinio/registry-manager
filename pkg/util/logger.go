package util

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const (
	// ALogFileName 访问日志
	ALogFileName = "./access.log"
	// LogFileName 应用日志
	LogFileName = "./app.log"
	// LogFileMaxSize 每个日志文件最大 MB
	LogFileMaxSize = 500
	// LogFileMaxBackups 保留日志文件个数
	LogFileMaxBackups = 10
	// LogFileMaxAge 保留日志最大天数
	LogFileMaxAge = 30
	// LogLevel 日志级别: -1Trace 0Debug 1Info 2Warn 3Error(默认) 4Fatal 5Panic 6NoLevel 7Off
	LogLevel = 0
)

var (
	ALog zerolog.Logger
	Log  zerolog.Logger
)

func init() {
	ALogger := &lumberjack.Logger{
		Filename:   ALogFileName,
		MaxSize:    LogFileMaxSize,
		MaxAge:     LogFileMaxAge,
		MaxBackups: LogFileMaxBackups,
		LocalTime:  true,
		Compress:   false,
	}
	Logger := &lumberjack.Logger{
		Filename:   LogFileName,
		MaxSize:    LogFileMaxSize,
		MaxAge:     LogFileMaxAge,
		MaxBackups: LogFileMaxBackups,
		LocalTime:  true,
		Compress:   false,
	}
	// 访问日志
	ALog = zerolog.New(zerolog.MultiLevelWriter(ALogger, os.Stdout)).With().Timestamp().Caller().Logger()
	// 应用日志
	Log = zerolog.New(zerolog.MultiLevelWriter(Logger, os.Stdout)).With().Timestamp().Caller().Logger()
}
