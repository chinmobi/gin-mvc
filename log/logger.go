// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package log

import (
	"os"

	"github.com/chinmobi/gin-mvc/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *zap.Logger = zap.L()
	SLogger *zap.SugaredLogger = zap.S()
)

func L() *zap.Logger {
	return Logger
}

func S() *zap.SugaredLogger {
	return SLogger
}

func setUpZap(config *config.Logger) error {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeTime  = zapcore.ISO8601TimeEncoder

	cores := []zapcore.Core{}

	if config.File.Enabled {
		hook := &lumberjack.Logger{
			Filename:   config.File.Filename,
			MaxSize:    int(config.File.MaxSize), // megabytes
			MaxBackups: int(config.File.MaxBackups),
			MaxAge:     int(config.File.MaxAge), //days
			Compress:   config.File.Compress,
		}

		fileLevel := mapLevel(config.File.Level)
		core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(hook), fileLevel)
		cores = append(cores, core)
	}

	if config.Console.Enabled {
		consoleLevel := mapLevel(config.Console.Level)
		core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), zapcore.Lock(os.Stdout), consoleLevel)
		cores = append(cores, core)
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core)
	Logger = logger

	SLogger = logger.Sugar()

	return nil
}

func mapLevel(str string) zapcore.Level {
	switch str {
		case "DEBUG":
			return zapcore.DebugLevel
		case "INFO":
			return zapcore.InfoLevel
		case "WARN":
			return zapcore.WarnLevel
		case "ERROR":
			return zapcore.ErrorLevel
		case "PANIC":
			return zapcore.PanicLevel
		case "FATAL":
			return zapcore.FatalLevel
		default:
			return zapcore.InfoLevel
	}
}
