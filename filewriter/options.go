package filewriter

import (
	"os"
	"path/filepath"
)

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

// fileWriter .
func (opts *Options) fileWriter() (*fileWriter, error) {
	fullpath, err := filepath.Abs(opts.FilePath)
	if err != nil {
		return nil, err
	}

	folderName, fileName := filepath.Split(fullpath)
	if err := os.MkdirAll(folderName, defaultFolderPermissions); err != nil {
		return nil, err
	}

	return &fileWriter{
		maxRolls:   opts.MaxRolls,
		folderName: folderName,
		fileName:   fileName,
		fullName:   fullpath,
	}, nil
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
