package log

import (
	"fmt"
	"go-projects-server/pkg/conf"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.SugaredLogger

// Init 初始化Logger
func Init(cfg *conf.Config) (err error) {
	var (
		l    zapcore.Level
		core zapcore.Core
	)
	writeSyncer := getLogWriter(
		cfg.Filename,   // 日志文件的位置
		cfg.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		cfg.MaxBackups, // 保留旧文件的最大个数
		cfg.MaxAge,     // 保留旧文件的最大天数
	)
	encoder := getEncoder()
	if err = l.UnmarshalText([]byte(cfg.Level)); err != nil {
		fmt.Println("zap init failed:", err)
		return
	}
	core = zapcore.NewCore(encoder, writeSyncer, l)
	lg = zap.New(core, zap.AddCaller()).Sugar()
	lg.Info("init log success")
	return
}

// getEncoder 设置zap编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 指定日志将写到哪里去，并使用Lumberjack进行日志切割归档
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(args ...interface{}) {
	lg.Debug(args)
}

func Info(args ...interface{}) {
	lg.Info(args)
}

func Warn(args ...interface{}) {
	lg.Warn(args)
}

func Error(args ...interface{}) {
	lg.Error(args)
}

func DPanic(args ...interface{}) {
	lg.DPanic(args)
}

func Panic(args ...interface{}) {
	lg.Panic(args)
}

func Fatal(args ...interface{}) {
	lg.Fatal(args)
}
