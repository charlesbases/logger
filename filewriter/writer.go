package filewriter

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
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

	// currentTime .
	currentTime time.Time
	// currentFileWriter os.File of log
	currentFileWriter *os.File
	// currentFileCreateAt log file create time
	currentFileCreateAt time.Time
	// currentFileExpireAt log file expire time
	currentFileExpireAt time.Time

	// lock zap 的日志输出流是线程安全的，此处的 lock 是防止零点时刻进行日志备份时，正确写入新的日志文件
	lock sync.Mutex
}

// Write .
func (fw *fileWriter) Write(p []byte) (int, error) {
	fw.lock.Lock()
	defer fw.lock.Unlock()

	fw.currentTime = time.Now()

	// needs to roll
	if !fw.currentFileExpireAt.IsZero() && !fw.currentTime.Before(fw.currentFileExpireAt) {
		if err := fw.rolling(); err != nil {
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
	}
	return err
}

// rolling .
func (fw *fileWriter) rolling() error {
	// close current file
	if err := fw.Close(); err != nil {
		return err
	}

	// rename
	if err := fw.rename(fw.currentFileCreateAt); err != nil {
		return err
	}

	go fw.tidy()
	return nil
}

// rename .
func (fw *fileWriter) rename(t time.Time) error {
	return os.Rename(fw.fullName, filepath.Join(fw.folderName, strings.Join([]string{fw.fileName, date(t)}, ".")))
}

// open .
func (fw *fileWriter) open() error {
	if err := os.MkdirAll(fw.folderName, defaultFolderPermissions); err != nil {
		return errors.Wrap(err, "mkdir folder")
	}

	fileInfo, err := os.Stat(fw.fullName)
	if err != nil {
		// 文件不存在，则直接创建新文件
		if os.IsNotExist(err) {
			return fw.create()
		}
		return err
	}

	// 是否为当天日志文件
	if fileInfo.ModTime().Day() != fw.currentTime.Day() ||
		fileInfo.ModTime().Year() != fw.currentTime.Year() ||
		fileInfo.ModTime().Month() != fw.currentTime.Month() {
		if err := fw.rename(fileInfo.ModTime()); err != nil {
			return err
		}
		go fw.tidy()
	}

	return fw.create()
}

// create .
func (fw *fileWriter) create() error {
	file, err := os.OpenFile(fw.fullName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, defaultFilePermissions)
	if err != nil {
		return errors.Wrap(err, "open file")
	}

	fw.currentFileWriter = file
	fw.currentFileCreateAt = midnight(time.Now())
	fw.currentFileExpireAt = midnight(fw.currentFileCreateAt, 1)
	return nil
}

// tidy .
func (fw *fileWriter) tidy() error {
	src, err := os.ReadDir(fw.folderName)
	if err != nil {
		return errors.Wrap(err, "open folder")
	}

	oldest := fw.oldest()

	for _, entry := range src {
		if !entry.IsDir() && len(entry.Name()) != len(fw.fileName) && strings.HasPrefix(entry.Name(), fw.fileName) {
			if suffix := filepath.Ext(entry.Name()); len(suffix) != 0 {
				suffix = suffix[1:]
				if t, err := time.Parse(defaultDateLayou, suffix); err == nil {
					if t.Before(oldest) {
						os.Remove(entry.Name())
					}
				}
			}
		}
	}
	return nil
}

// oldest return fileWriter.currentTime - fileWriter.maxRolls * time.Hour * 24
func (fw *fileWriter) oldest() time.Time {
	return midnight(fw.currentTime, -fw.maxRolls)
}

// New .
func New(opts ...func(o *options)) *fileWriter {
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

// date .
func date(t time.Time) string {
	return t.Format(defaultDateLayou)
}

// midnight 零点时间
func midnight(t time.Time, days ...int) time.Time {
	var day int
	if len(days) != 0 {
		day = days[0]
	}
	return time.Date(t.Year(), t.Month(), t.Day()+day, 0, 0, 0, 0, t.Location())
}
