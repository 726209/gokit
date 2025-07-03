package logger

import (
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
	"go.uber.org/zap"
	"os"
	"strconv"
)

var usePterm bool
var globalLogLevel LogLevel

// Config 日志配置项
type Config struct {
	Identifier  string       // 唯一码（创建目录）
	Level       string       // 日志等级
	PrettyPrint bool         // 是否美化输出（使用 pterm）
	OutputPath  string       // 日志文件输出路径，若为空仅控制台输出
	ZapOptions  []zap.Option // 可扩展的 zap 配置项
}

func (c *Config) merge(option Config) {
	if option.Level != "" {
		c.Level = option.Level
	}
	c.PrettyPrint = option.PrettyPrint || c.PrettyPrint
	if option.OutputPath != "" {
		c.OutputPath = option.OutputPath
	}
	if len(option.ZapOptions) > 0 {
		c.ZapOptions = append(c.ZapOptions, option.ZapOptions...)
	}
}

// InitLogger 初始化日志系统
// 支持可选参数，默认配置 + 可选 Config 重载
// Level: 日志级别，如 "info"、"warn"、"error"、"debug"
// PrettyPrint: 是否使用控制台彩色日志
// OutputPath: 日志写入文件（使用 lumberjack 滚动写入）
// If not provided, defaults to "info" and pretty=false.
// Example:
//
//	logger.InitLogger(logger.Config{
//		Level:       "info",
//		PrettyPrint: true,
//	})
//	defer logger.Sync()
func InitLogger(options ...Config) {
	// 默认配置
	config := Config{
		Identifier:  "com.yourcompany.yourapp",
		Level:       "Info",
		PrettyPrint: false,
		OutputPath:  "",
		ZapOptions:  nil,
	}
	// 读取环境变量
	_ = godotenv.Load()
	pretty, _ := strconv.ParseBool(os.Getenv("LOG_PRETTY"))
	config.merge(Config{
		Level:       os.Getenv("LOG_LEVEL"),
		PrettyPrint: pretty,
		OutputPath:  os.Getenv("LOG_PATH"),
	})
	// 依次合并每个 Config
	for _, cfg := range options {
		config.merge(cfg)
	}
	usePterm = config.PrettyPrint
	globalLogLevel = ParseLevel(config.Level)
	built(config)
}

func Sync() {
	if zapLogger != nil {
		_ = zapLogger.Sync()
	}
}

func Pack(args ...any) []pterm.LoggerArgument {
	//arguments := pterm.DefaultLogger.Args(args...)
	//for i := range arguments {
	//	switch v := arguments[i].Value.(type) {
	//	case time.Duration:
	//		arguments[i].Value = v.String() // 转为 "1s"、"150ms" 等格式
	//		arguments[i].Value = v.Nanoseconds()
	//	}
	//}
	//return arguments
	return pterm.DefaultLogger.Args(args...)
}

func PackMap(m map[string]any) []pterm.LoggerArgument {
	//for k, v := range m {
	//	switch val := v.(type) {
	//	case time.Duration:
	//		m[k] = val.String()
	//	}
	//}
	return pterm.DefaultLogger.ArgsFromMap(m)
}

func Trace(msg string, args ...[]pterm.LoggerArgument) { log(TraceLevel, msg, args...) }
func Tracef(template string, args ...interface{})      { logf(TraceLevel, template, args...) }

func Debug(msg string, args ...[]pterm.LoggerArgument) { log(DebugLevel, msg, args...) }
func Debugf(template string, args ...interface{})      { logf(DebugLevel, template, args...) }

func Info(msg string, args ...[]pterm.LoggerArgument) { log(InfoLevel, msg, args...) }
func Infof(template string, args ...interface{})      { logf(InfoLevel, template, args...) }

func Warn(msg string, args ...[]pterm.LoggerArgument) { log(WarnLevel, msg, args...) }
func Warnf(template string, args ...interface{})      { logf(WarnLevel, template, args...) }

func Error(msg string, args ...[]pterm.LoggerArgument) { log(ErrorLevel, msg, args...) }
func Errorf(template string, args ...interface{})      { logf(ErrorLevel, template, args...) }

func Fatal(msg string, args ...[]pterm.LoggerArgument) { log(FatalLevel, msg, args...) }
func Fatalf(template string, args ...interface{})      { logf(FatalLevel, template, args...) }

func Print(msg string, args ...[]pterm.LoggerArgument) { log(PrintLevel, msg, args...) }
func Printf(template string, args ...interface{})      { logf(PrintLevel, template, args...) }

//func TraceSQL(ctx context.Context, sql string, elapsed time.Duration, rows int64, err error) {
//	msg := "SQL: %s | %v | rows: %d"
//	if err != nil {
//		Error("SQL ERROR: "+sql, err, rows)
//	} else if elapsed > 200*time.Millisecond {
//		Warn("SLOW QUERY: "+sql, elapsed, rows)
//	} else {
//		Debug(msg, sql, elapsed, rows)
//	}
//}
