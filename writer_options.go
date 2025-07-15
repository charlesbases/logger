package logger

// WriterOption 定义配置文件日志写入器的函数类型。
// 用于通过函数式选项模式定制日志文件的各种参数。
type WriterOption func(w *fileWriter) ()

// WithFileDir 设置日志文件保存的目录选项。
//
// 参数:
//   dir - 日志文件存储目录的绝对或相对路径
//
// 示例：
//   WithFileDir("/var/log/myapp")
func WithFileDir(dir string) WriterOption {
	return func(w *fileWriter) {
		w.filedir = dir
	}
}

// WithFileName 设置日志文件名称选项。
//
// 参数:
//   name - 日志文件的名称
//
// 示例：
//   WithFileName("app.log")
func WithFileName(name string) WriterOption {
	return func(w *fileWriter) {
		w.filename = name
	}
}

// WithMaxAge 设置日志文件的最大保留天数选项。
//
// 参数:
//   maxAge - 日志文件保留的最大天数。超过此天数的日志文件会被删除。
//
// 特殊值：
//   0 - 永远保留所有日志文件
//
// 示例：
//   WithMaxAge(60) // 保留60天的日志
func WithMaxAge(maxAge int) WriterOption {
	return func(w *fileWriter) {
		w.maxAge = maxAge
	}
}

// WithMaxSize 设置单个日志文件的最大大小选项。
//
// 参数:
//   maxSize - 单个日志文件的最大大小（单位：MB）
//
// 当日志文件达到此大小时，会自动进行轮转创建新文件。
//
// 示例：
//   WithMaxSize(256) // 每个日志文件最大256MB
func WithMaxSize(maxSize int) WriterOption {
	return func(w *fileWriter) {
		w.maxSize = maxSize
	}
}

// WithCompress 设置是否压缩旧日志文件选项。
//
// 参数:
//   b - 是否启用压缩：
//       true - 轮转后的旧日志文件会被压缩为 .gz 格式
//       false - 不压缩旧日志文件
//
// 注意：压缩操作是异步进行的，不会阻塞当前日志写入。
//
// 示例：
//   WithCompress(true) // 启用日志压缩
func WithCompress(b bool) WriterOption {
	return func(w *fileWriter) {
		w.compress = b
	}
}
