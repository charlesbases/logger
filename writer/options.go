package writer

import (
	"os"
	"path/filepath"
)

const (
	// defaultMaxRolls 日志保留时间
	defaultMaxRolls = 7
	// defaultFilePath default file path
	defaultFilePath = "./log/log"
)

// Option .
type Option interface {
	apply(o *options)
}

// optfunc .
type optfunc func(o *options) ()

// apply .
func (f optfunc) apply(o *options) () {
	f(o)
}

// FilePath .
func FilePath(n string) Option {
	return optfunc(
		func(o *options) {
			o.filePath = n
		},
	)
}

// MaxRolls .
func MaxRolls(n int) Option {
	return optfunc(
		func(o *options) {
			o.maxRolls = n
		},
	)
}

// options .
type options struct {
	// filePath 日志文件路径
	filePath string
	// maxRolls 日志文件保留天数
	maxRolls int
}

// fileWriter .
func (opts *options) fileWriter() (*fileWriter, error) {
	fullpath, err := filepath.Abs(opts.filePath)
	if err != nil {
		return nil, err
	}

	folderName, fileName := filepath.Split(fullpath)
	if err := os.MkdirAll(folderName, defaultFolderPermissions); err != nil {
		return nil, err
	}

	return &fileWriter{
		maxRolls:   opts.maxRolls,
		folderName: folderName,
		fileName:   fileName,
		fullName:   fullpath,
	}, nil
}

// configuration .
func configuration(opts ...Option) *options {
	var options = new(options)
	for _, opt := range opts {
		opt.apply(options)
	}

	if len(options.filePath) == 0 {
		options.filePath = defaultFilePath
	}
	if options.maxRolls < 1 {
		options.maxRolls = defaultMaxRolls
	}

	return options
}
