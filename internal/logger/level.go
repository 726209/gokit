package logger

import (
	"github.com/pterm/pterm"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"

	"strings"
)

type LogLevel string

const (
	NoneLevel  LogLevel = "none"
	TraceLevel LogLevel = "trace"
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
	PrintLevel LogLevel = "print"
)

func ParseLevel(str string) LogLevel {
	switch strings.ToLower(str) {
	case "none":
		return NoneLevel
	case "trace":
		return TraceLevel
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "print":
		return PrintLevel
	default: // info
		return InfoLevel
	}
}

// Zap 转换为 zapcore.Level
func (l LogLevel) Zap() zapcore.Level {
	switch l {
	case DebugLevel:
		return zapcore.DebugLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func GormLogLevel() gormlogger.LogLevel {
	return globalLogLevel.Gorm()
}

func (l LogLevel) Gorm() gormlogger.LogLevel {
	switch l {
	case NoneLevel:
		return gormlogger.Silent
	case ErrorLevel, FatalLevel:
		return gormlogger.Error
	case WarnLevel:
		return gormlogger.Warn
	default:
		return gormlogger.Info // gorm 无 Debug，用 Info 替代
	}
}

func (l LogLevel) GormString() string {
	level := l.Gorm()
	switch level {
	case gormlogger.Silent:
		return "Silent"
	case gormlogger.Error:
		return "Error"
	case gormlogger.Warn:
		return "Warn"
	default:
		return "Info"
	}
}

// Pterm 转换为 pterm.Level
func (l LogLevel) Pterm() pterm.LogLevel {
	switch l {
	case NoneLevel:
		return pterm.LogLevelDisabled
	case TraceLevel:
		return pterm.LogLevelTrace
	case DebugLevel:
		return pterm.LogLevelDebug
	case WarnLevel:
		return pterm.LogLevelWarn
	case ErrorLevel:
		return pterm.LogLevelError
	case FatalLevel:
		return pterm.LogLevelFatal
	case PrintLevel:
		return pterm.LogLevelPrint
	default:
		return pterm.LogLevelInfo
	}
}
