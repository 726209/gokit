package repository

import (
	"context"
	"github.com/726209/gokit/logger"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	LogLevel      gormlogger.LogLevel // 日志级别
	SlowThreshold time.Duration       // 慢查询阈值
}

// NewGormLogger /*
// Example
//
//	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//			Logger: &GormLogger{
//				LogLevel:      logger.GormLogLevel(),
//				SlowThreshold: 200 * time.Millisecond,
//			},
//	})
func NewGormLogger() gormlogger.Interface {
	return &GormLogger{
		LogLevel:      logger.GormLogLevel(),
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	l.LogLevel = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

func (l *GormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

const maxLogLength = 500 // 最多打印500个字符

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel == gormlogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()

	sql = truncate(sql, maxLogLength)

	switch {
	case err != nil && l.LogLevel >= gormlogger.Error:
		logger.Errorf("[SQL] %s | error: %v | rows=%d", sql, err, rows)
	case elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		logger.Warnf("[SLOW] %s | duration: %s | rows=%d", sql, elapsed, rows)
	case l.LogLevel >= gormlogger.Info:
		logger.Debugf("[SQL] %s | duration: %s | rows=%d", sql, elapsed, rows)
	}
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max] + "...[truncated]"
	}
	return s
}
