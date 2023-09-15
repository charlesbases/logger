package filewriter

import "path/filepath"

// options .
type options struct {
	// output 日志文件路径
	output string
	// maxrolls 日志文件保留天数
	maxrolls int
}

// Path .
func Path(file string) func(o *options) {
	return func(o *options) {
		if len(file) != 0 {
			o.output = file
		}
	}
}

// MaxRolls .
func MaxRolls(days int) func(o *options) {
	return func(o *options) {
		if days != 0 {
			o.maxrolls = days
		}
	}
}

// option .
func option(opts ...func(o *options)) *fileWriter {
	options := &options{
		output:   defaultPath,
		maxrolls: defaultMaxRolls,
	}
	for _, opt := range opts {
		opt(options)
	}

	fullpath, _ := filepath.Abs(options.output)
	folderName, fileName := filepath.Split(fullpath)

	return &fileWriter{
		maxRolls:   options.maxrolls,
		folderName: folderName,
		fileName:   fileName,
		fullName:   fullpath,
	}
}
