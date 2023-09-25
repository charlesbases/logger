package filewriter

const (
	// defaultMaxRolls 日志保留时间
	defaultMaxRolls = 7
	// defaultFilePath default file path
	defaultFilePath = "./logs/log"
)

// Options .
type Options struct {
	// FilePath 日志文件路径
	FilePath string
	// MaxRolls 日志文件保留天数
	MaxRolls int
}

// configuration .
func configuration(opts ...func(o *Options)) *Options {
	var options = new(Options)
	for _, opt := range opts {
		opt(options)
	}

	if len(options.FilePath) == 0 {
		options.FilePath = defaultFilePath
	}
	if options.MaxRolls < 1 {
		options.MaxRolls = defaultMaxRolls
	}

	return options
}
