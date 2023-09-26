package filewriter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	// defaultFilePermissions default file permissions
	defaultFilePermissions = 0666
	// defaultFolderPermissions default folder prmissions
	defaultFolderPermissions = 0775
	// defaultSuffixFormat the date suffix of the log file
	defaultSuffixFormat = "2006-01-02"
	// defaultFormatLayou format layou
	defaultFormatLayou = "2006-01-02 15:04:05.000"
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

	// currentTime now
	currentTime time.Time
	// currentFileWriter os.File of log
	currentFileWriter *os.File
	// currentFileCreateAt lof file creation date
	currentFileCreateAt time.Time
	// currentFileExpireAt lof file expiration date
	currentFileExpireAt time.Time

	lock sync.Mutex
}

// Write .
func (fw *fileWriter) Write(p []byte) (int, error) {
	fw.lock.Lock()
	defer fw.lock.Unlock()

	n, err := fw.write(p)
	if err != nil {
		stderr(err)
	}
	return n, nil
}

// write .
func (fw *fileWriter) write(p []byte) (int, error) {
	fw.currentTime = time.Now()

	// needs to roll
	if fw.currentFileWriter != nil && !fw.currentTime.Before(fw.currentFileExpireAt) {
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

	return nil
}

// rename .
func (fw *fileWriter) rename(timeSuffix time.Time) error {
	return os.Rename(fw.fullName, filepath.Join(fw.folderName, strings.Join([]string{fw.fileName, timeString(timeSuffix)}, ".")))
}

// open .
func (fw *fileWriter) open() error {
	fileInfo, err := os.Stat(fw.fullName)
	if err != nil {
		// 文件不存在，则直接创建新文件
		if os.IsNotExist(err) {
			return fw.create()
		}
		return err
	}

	// 是否为当天日志文件
	if !timeMidnight(fw.currentTime, 0).Equal(timeMidnight(fileInfo.ModTime(), 0)) {
		if err := fw.rename(fileInfo.ModTime()); err != nil {
			return err
		}
	}

	return fw.create()
}

// create .
func (fw *fileWriter) create() error {
	go fw.tidy()

	file, err := os.OpenFile(fw.fullName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, defaultFilePermissions)
	if err != nil {
		return err
	}

	fw.currentFileWriter = file
	fw.currentFileCreateAt = fw.currentTime
	fw.currentFileExpireAt = timeMidnight(fw.currentTime, 1)
	return nil
}

// tidy .
func (fw *fileWriter) tidy() error {
	src, err := os.ReadDir(fw.folderName)
	if err != nil {
		return err
	}
	if len(src) == 0 {
		return nil
	}

	oldest := timeMidnight(fw.currentTime, -fw.maxRolls)

	for _, entry := range src {
		if !entry.IsDir() && len(entry.Name()) != len(fw.fileName) && strings.HasPrefix(entry.Name(), fw.fileName) {
			if suffix := filepath.Ext(entry.Name()); len(suffix) != 0 {
				suffix = suffix[1:]
				if t, err := time.ParseInLocation(defaultSuffixFormat, suffix, fw.currentTime.Location()); err == nil {
					if t.Before(oldest) {
						os.Remove(filepath.Join(fw.folderName, entry.Name()))
					}
				}
			}
		}
	}
	return nil
}

// New .
func New(opts ...func(o *Options)) io.Writer {
	fileWriter, err := configuration(opts...).fileWriter()
	if err != nil {
		stderr(err)
		return nil
	}
	return fileWriter
}

// timeString .
func timeString(t time.Time) string {
	return t.Format(defaultSuffixFormat)
}

// timeMidnight 零点时间
func timeMidnight(t time.Time, offset int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+offset, 0, 0, 0, 0, t.Location())
}

// stderr .
func stderr(err error) {
	if err != nil {
		fmt.Printf("[%s] \033[31mERR\033[0m %s %v\n", time.Now().Format(defaultFormatLayou), caller(), err)
	}
}

// caller .
func caller() string {
	if _, file, line, ok := runtime.Caller(1); ok {
		return fmt.Sprintf(`%s/%s:%d`, filepath.Base(filepath.Dir(file)), filepath.Base(file), line)
	}
	return "undefined"
}
