package logger

import (
	"io"
	"strings"

	"go.uber.org/zap/zapcore"
)

// Option 是用于配置 zapLogger 的函数类型
//
// 这种函数式选项模式提供了灵活且可扩展的配置方式，
// 允许在不破坏现有代码的情况下添加新配置项。
type Option func(l *zapLogger) ()

// WithName 设置日志记录器的名称选项
//
// 参数:
//   name - 日志记录器名称（通常为应用或模块名）
//
// 示例:
//   WithName("app")
func WithName(name string) Option {
	return func(l *zapLogger) {
		l.name = name
	}
}

// WithWriter 设置日志输出目标选项
//
// 参数:
//   ws... - 一个或多个实现了 io.Writer 接口的输出目标
//
// 特性:
//   - 支持同时输出到多个目标
//   - 内置线程安全处理
//   - 支持任何实现 io.Writer 接口的类型
//
// 示例:
//   // 输出到控制台和文件
//   WithWriter(os.Stdout, fileWriter)
func WithWriter(ws ...io.Writer) Option {
	return func(l *zapLogger) {
		l.writers = ws
	}
}

// WithColorful 设置控制台输出是否启用彩色选项
//
// 参数:
//   b - 是否启用彩色输出:
//        true  - 在支持颜色的终端显示彩色日志级别
//        false - 始终使用单色输出
func WithColorful(b bool) Option {
	return func(l *zapLogger) {
		l.colorful = b
	}
}

// WithMinLevel 设置最低日志级别选项
//
// 参数:
//   level - 字符串表示的日志级别（不区分大小写）
//           支持的值: "debug", "info", "warn", "error", "panic", "fatal"
//
// 默认行为:
//   如果提供无效级别字符串，默认使用 DebugLevel
//
// 示例:
//   WithMinLevel("info") // 只记录 zapcore.InfoLevel 及以上级别的日志
func WithMinLevel(level string) Option {
	return func(l *zapLogger) {
		switch strings.ToLower(level) {
		case "debug":
			l.minLevel = zapcore.DebugLevel
		case "info":
			l.minLevel = zapcore.InfoLevel
		case "warn":
			l.minLevel = zapcore.WarnLevel
		case "error":
			l.minLevel = zapcore.ErrorLevel
		case "panic":
			l.minLevel = zapcore.PanicLevel
		case "fatal":
			l.minLevel = zapcore.FatalLevel
		default:
			l.minLevel = zapcore.DebugLevel
		}
	}
}
