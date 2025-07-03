package logger

import (
	"encoding/json"
	"fmt"
	"github.com/726209/gokit"
	"github.com/pterm/pterm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
)

var zapLogger *zap.Logger

func built(option Config) {
	pterm.FgGreen.Println(">>> 正在初始化，请稍后...")
	// 创建核心组件，多个输出合并
	core := zapcore.NewTee(builtFileCore(option), builtConsoleCore(option))
	// 返回 logger（带 caller 信息）
	// 在日志等级 ≥ Error（包括 Error、DPanic、Panic、Fatal） 时，自动添加 调用堆栈信息
	logger := zap.New(core,
		zap.AddCaller(),
		//zap.AddCallerSkip(3), // 跳过1层
		zap.AddStacktrace(zapcore.ErrorLevel))
	// 可选：设置为全局默认 logger
	zap.ReplaceGlobals(logger)
	zapLogger = logger
	printBootLogInfo(option)
}

func getLogFilepath(option Config) string {
	file := option.OutputPath
	if file == "" {
		file = gokit.DefaultLogPathWithName(option.Identifier, "app.log")
	}
	_ = os.MkdirAll(filepath.Dir(file), 0755)
	return file
}

// 配置 lumberjack 滚动日志
func builtFileCore(option Config) zapcore.Core {
	// 1. 文件输出：配置 lumberjack 滚动日志
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   getLogFilepath(option), // 日志文件路径
		MaxSize:    10,                     // 每个日志文件最大 MB
		MaxBackups: 5,                      // 最多保留旧文件数量
		MaxAge:     28,                     // 最多保留天数
		Compress:   true,                   // 是否压缩
	})
	// 2. 设置日志级别和编码器
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeDuration = zapcore.StringDurationEncoder
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(cfg)
	return zapcore.NewCore(encoder, writer, globalLogLevel.Zap())
}

// 配置控制台输出
func builtConsoleCore(option Config) zapcore.Core {
	if !option.PrettyPrint {
		cfg := zap.NewDevelopmentEncoderConfig()
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder := zapcore.NewConsoleEncoder(cfg)
		return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), globalLogLevel.Zap())
	} else {
		builtPtermLogger()
		// 如果 pretty 为 true，控制台输出交给 pterm，不初始化 consoleCore
		return zapcore.NewNopCore()
	}
}

// 如果启用了 pretty，控制台彩色日志
func builtPtermLogger() {
	level := globalLogLevel.Pterm()
	// Define a new style for the "priority" key.
	priorityStyle := map[string]pterm.Style{
		"priority": *pterm.NewStyle(pterm.FgRed),
	}
	pterm.PrintDebugMessages = level <= pterm.LogLevelDebug // 启用 debug 输出
	width := pterm.GetTerminalWidth()
	if width <= 0 {
		width = 120 // fallback 宽度
	}
	custom := pterm.DefaultLogger.
		WithMaxWidth(width).
		WithLevel(level). // 设置日志级别
		WithTime(true).   // 显示时间
		//WithTimeFormat("15:04:05").      // 时间格式
		//WithCaller(pterm.PrintDebugMessages). // 显示调用者
		WithCallerOffset(1).
		//WithFormatter(pterm.LogFormatterJSON).
		WithKeyStyles(priorityStyle).WithFormatter(pterm.LogFormatterJSON)

	// Define a new style for the "foo" key.
	fooStyle := *pterm.NewStyle(pterm.FgGreen)

	// Append the new style to the existing ones.
	custom.AppendKeyStyle("devt", fooStyle)
	pterm.DefaultLogger = *custom
}

// 定义样式：浅蓝前景 + 灰色背景 + 加粗
// var bootInfo = pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold)
// bootInfo.Println("🔧 CLI 启动配置：")
// bootInfo.Printfln("🚀 当前版本:   %s", version.Version)
func printBootLogInfo(option Config) {
	labelStyle := pterm.NewStyle(pterm.FgLightCyan, pterm.Bold)
	textStyle := pterm.NewStyle(pterm.FgWhite)
	detail := fmt.Sprintf("Pterm(%s,debug=%t) Zap(%s) Gorm(%s)",
		pterm.DefaultLogger.Level.String(),
		pterm.PrintDebugMessages,
		zap.L().Level().String(),
		globalLogLevel.GormString(),
	)
	lines := []string{
		pterm.Info.Sprintf("👋 你好，欢迎使用 %s！", option.Identifier),
		fmt.Sprintf("%s %s", labelStyle.Sprint("🚀 当前版本: "), textStyle.Sprint("version.Version")),
		fmt.Sprintf("%s %s", labelStyle.Sprint("📘 日志配置:"), textStyle.Sprintf("--level=%s --pretty=%t %s", globalLogLevel, option.PrettyPrint, detail)),
		fmt.Sprintf("%s %s", labelStyle.Sprint("📁 日志文件:"), textStyle.Sprintf("` tail -f \"%s\" `", getLogFilepath(option))),
	}
	pterm.DefaultBox.
		WithTitle("🔧 启动配置").
		WithTitleTopLeft().
		WithLeftPadding(2).
		WithRightPadding(2).Println(pterm.LightWhite(strings.Join(lines, "\n")))
}

// 通用日志方法：支持结构化和格式化
func write(level LogLevel, msg string, skip int, args ...[]pterm.LoggerArgument) {
	// .Check(level, msg) 预检查日志是否需要记录，比如当前日志级别为 Info，那么只有 Info 及以上的才会继续执行。
	// 相比 logger.Info(msg, zap.String("args", ...))，这种写法性能更高，避免了在日志级别不足时仍构造字段的开销。
	if entry := zap.L().
		WithOptions(zap.AddCallerSkip(skip)).
		Check(level.Zap(), msg); entry != nil {
		if len(args) > 0 {
			entry.Write(zap.Any("extra", argsToFlatMap(args...)))
		} else {
			entry.Write()
		}
	}
	// pterm 日志打印逻辑
	if usePterm {
		pLevel := level.Pterm()
		switch pLevel {
		case pterm.LogLevelTrace:
			pterm.DefaultLogger.Trace(msg, args...)
		case pterm.LogLevelDebug:
			pterm.DefaultLogger.Debug(msg, args...)
		case pterm.LogLevelInfo:
			pterm.DefaultLogger.Info(msg, args...)
		case pterm.LogLevelWarn:
			pterm.DefaultLogger.Warn(msg, args...)
		case pterm.LogLevelError:
			pterm.DefaultLogger.Error(msg, args...)
		case pterm.LogLevelFatal:
			pterm.DefaultLogger.Fatal(msg, args...)
		case pterm.LogLevelPrint:
			pterm.DefaultLogger.Print(msg, args...)
		default:
			// 默认 LogLevelDisabled
		}
	}
}

// log 统一结构化输出 日志，根据 Pretty 控制是否 pterm 打印
func log(level LogLevel, msg string, args ...[]pterm.LoggerArgument) {
	write(level, msg, 3, args...) // skip=1 调用者是 log()
}

// log 统一格式化输出 日志，根据 Pretty 控制是否 pterm 打印
func logf(level LogLevel, template string, args ...interface{}) {
	write(level, fmt.Sprintf(template, args...), 3) // skip=2 调用者是 logf() → write()
}

// argsToMap 将 LoggerArgument 多组参数转换为扁平的 map[string]any
func argsToFlatMap(args ...[]pterm.LoggerArgument) map[string]any {
	// 转换为 map
	flatmap := make(map[string]any)
	for _, group := range args {
		for _, arg := range group {
			flatmap[arg.Key] = arg.Value
		}
	}
	return flatmap
}

// argsToJSON 将 LoggerArgument 数组转换为 JSON 字符串
func argsToJSON(args ...[]pterm.LoggerArgument) string {
	// 转换为 map
	flatmap := argsToFlatMap(args...)
	if len(flatmap) == 0 {
		return ""
	}

	// 转为 JSON 字符串
	jsonStr, err := json.Marshal(flatmap)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}
