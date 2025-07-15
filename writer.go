package logger

import (
	"io"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

type fileWriter struct {
	filedir   string
	filename  string
	maxAge    int
	maxSize   int
	compress  bool
	localtime bool
}

// NewFileWriter 创建并返回一个实现了 io.Writer 接口的文件日志写入器。
//
// 该写入器使用 lumberjack 库实现日志轮转功能，支持按文件大小和时间自动轮转日志文件。
// 提供了多个 WriterOption 函数用于自定义日志文件的行为参数。
//
// 默认参数：
//   - 日志目录: "./logs"
//   - 文件名: "log.log"
//   - 最大保留天数: 0
//   - 单个文件最大大小: 128MB
//   - 使用本地时间: true
//   - 压缩旧日志: false
//
// 示例：
//   writer := NewFileWriter(
//       WithFileDir("./logs"),
//       WithFileName("app.log"),
//       WithMaxAge(7),
//       WithMaxSize(512),
//       WithCompress(true),
//   )
//
func NewFileWriter(opts ...WriterOption) io.Writer {
	options := &fileWriter{
		filedir:   "./logs",
		filename:  "log.log",
		maxSize:   128,
		localtime: true,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &lumberjack.Logger{
		Filename:  filepath.Join(options.filedir, options.filename),
		MaxSize:   options.maxSize,
		MaxAge:    options.maxAge,
		LocalTime: options.localtime,
		Compress:  options.compress,
	}
}
