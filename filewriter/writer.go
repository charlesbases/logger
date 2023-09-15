package filewriter

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	// defaultMaxRolls 日志保留时间
	defaultMaxRolls = 7
	// defaultPath default file path
	defaultPath = "./logs/log"
	// defaultFilePermissions default file permissions
	defaultFilePermissions = 0666
	// defaultFolderPermissions default folder prmissions
	defaultFolderPermissions = 0775
	// defaultDateLayou format layou
	defaultDateLayou = "2006-01-02"
)

// fileWriter is used to write to a file.
type fileWriter struct {
	maxRolls int

	// folderName log folder name
	folderName string
	// fileName log file name
	fileName string
	// fullName log file name of abs
	fullName string

	// currentFileWriter os.File of log
	currentFileWriter *os.File
	// currentFileExpireAt log file expire time
	currentFileExpireAt int64

	// lock zap 的日志输出流是线程安全的，此处的 lock 是防止零点时刻进行日志备份时，正确写入新的日志文件
	lock sync.Mutex
}

// New .
func New(opts ...func(o *options)) *fileWriter {
	return option(opts...)
}

// Write .
func (fw *fileWriter) Write(p []byte) (int, error) {
	t := time.Now()

	fw.lock.Lock()
	defer fw.lock.Unlock()

	// needs to roll
	if fw.currentFileExpireAt != 0 && fw.currentFileExpireAt <= t.Unix() {
		if err := fw.rolling(t); err != nil {
			return 0, err
		}
	}

	// needs to create
	if fw.currentFileWriter == nil {
		if err := fw.open(); err != nil {
			return 0, err
		}
	}
	return fw.currentFileWriter.Write(p)
}

// Close .
func (fw *fileWriter) Close() (err error) {
	if fw.currentFileWriter != nil {
		err = fw.currentFileWriter.Close()
		fw.currentFileWriter = nil
		fw.currentFileExpireAt = 0
	}
	return err
}

// open .
func (fw *fileWriter) open() error {
	t := time.Now()

	if err := os.MkdirAll(fw.folderName, defaultFolderPermissions); err != nil {
		return err
	}

	fileInfo, err := os.Stat(fw.fullName)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return fw.create(t)
	}

	// 当前已存在的日志文件为旧文件
	if next(fileInfo.ModTime(), 1) <= t.Unix() {
		if err := fw.rename(t); err != nil {
			return err
		}

		go fw.tidy()
	}

	return fw.create(t)
}

// create .
func (fw *fileWriter) create(t time.Time) error {
	file, err := os.OpenFile(fw.fullName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, defaultFilePermissions)
	if err != nil {
		return err
	}

	fw.currentFileWriter = file
	fw.currentFileExpireAt = next(t, 1)
	return nil
}

// rolling .
func (fw *fileWriter) rolling(t time.Time) error {
	// close current file
	if err := fw.Close(); err != nil {
		return err
	}

	// rename
	if err := fw.rename(t); err != nil {
		return err
	}

	go fw.tidy()
	return nil
}

// rename .
func (fw *fileWriter) rename(t time.Time) error {
	return os.Rename(fw.fullName, filepath.Join(fw.folderName, strings.Join([]string{fw.fileName, suffix(t)}, ".")))
}

// oldest 根据 fileWriter.maxRolls 获取最旧的日期
func (fw *fileWriter) oldest() int64 {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day()-fw.maxRolls, 0, 0, 0, 0, t.Location()).Unix()
}

// tidy remove old log
func (fw *fileWriter) tidy() error {
	src, err := os.ReadDir(fw.folderName)
	if err != nil {
		return err
	}

	oldest := fw.oldest()

	for _, entry := range src {
		// old log
		if !entry.IsDir() && len(entry.Name()) != len(fw.fileName) && strings.HasPrefix(entry.Name(), fw.fileName) {
			if suffix := filepath.Ext(entry.Name()); len(suffix) != 0 {
				suffix = suffix[1:]
				if t, err := time.Parse(defaultDateLayou, suffix); err == nil {
					if next(t, 0) < oldest {
						os.Remove(entry.Name())
					}
				}
			}
		}
	}
	return nil
}

// suffix .
func suffix(t time.Time) string {
	return t.Format(defaultDateLayou)
}

// next .
func next(t time.Time, days int) int64 {
	return time.Date(t.Year(), t.Month(), t.Day()+days, 0, 0, 0, 0, t.Location()).Unix()
}
