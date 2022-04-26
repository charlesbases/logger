package filewriter

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	// DefaultMaxRolls 日志保留时间
	DefaultMaxRolls = 7
	// DefaultFilePath default file path
	DefaultFilePath = "./log/log.log"
)

type fileWriter struct {
	fileDir    string
	fileName   string
	filePath   string
	filePrefix string

	opts      *Options
	writer    io.WriteCloser
	current   *fileInfo
	fileStack *fileStack

	lock    sync.Mutex
	closing chan struct{}
}

// Options .
type Options struct {
	// TTL 日志文件保留天数
	TTL int
	// FilePath 日志文件路径
	FilePath string
}

// defaultOptions .
func defaultOptions() *Options {
	return &Options{
		TTL:      DefaultMaxRolls,
		FilePath: DefaultFilePath,
	}
}

type Option func(o *Options)

// TTL 日志文件保留天数
func TTL(ttl int) Option {
	return func(o *Options) {
		o.TTL = ttl
	}
}

// FilePath .
func FilePath(fpath string) Option {
	return func(o *Options) {
		o.FilePath = fpath
	}
}

// New .
func New(options ...Option) *fileWriter {
	var opts = defaultOptions()
	for _, o := range options {
		o(opts)
	}

	var fw = &fileWriter{
		opts: opts,
		fileStack: &fileStack{
			files: make([]*fileInfo, opts.TTL-1),
			cap:   opts.TTL - 1,
		},
	}

	if err := fw.repo(opts.FilePath); err != nil {
		panic(err)
	}
	return fw
}

// Write .
func (fw *fileWriter) Write(p []byte) (n int, err error) {
	fw.lock.Lock()
	if fw.isNotNil() {
		n, err = fw.writer.Write(p)
	}
	fw.lock.Unlock()
	return n, err
}

// Close close fileWriter
func (fw *fileWriter) Close() {
	fw.closing <- struct{}{}
}

// closeCurrentFile close current file
func (fw *fileWriter) closeCurrentFile() {
	if fw.current != nil && fw.writer != nil {
		fw.writer.Close()

		fw.writer = nil
		fw.current = nil
	}
}

// isNotNil .
func (fw *fileWriter) isNotNil() bool {
	return fw.current != nil && fw.writer != nil
}

func (fw *fileWriter) repo(filename string) error {
	abs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	fw.filePath = abs
	fw.fileDir, fw.fileName = filepath.Split(abs)
	fw.filePrefix = fw.filePathJoin(fw.fileName, ".")

	return fw.loading()
}

// loading load fileStack.files
func (fw *fileWriter) loading() error {
loading:
	files, err := ioutil.ReadDir(fw.fileDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(fw.fileDir, 0777); err != nil {
				return err
			}
			goto loading
		}
		return err
	}

	current := newFileDate(fw.opts.TTL)

	// 统计历史日志文件
	if len(files) != 0 {
		list := make([]*fileInfo, 0, len(files))
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			switch {
			case file.Name() == fw.fileName && file.Size() != 0:
				// 日志文件不为空时
				// 查看是否为当天日志文件
				// 如果日期为当天，则追加写入
				// 如果日期不为当天，修改文件名后缀，当天日志记录在新文件中
				fileDate := newFileDateWithTime(file.ModTime(), fw.opts.TTL)

				switch {
				case current.createAt.Equal(fileDate.createAt):
					// nothing
				case current.createAt.After(fileDate.expireAt) || current.createAt.Equal(fileDate.expireAt):
					// remove
					os.Remove(fw.filePath)
				case current.createAt.After(fileDate.createAt) && current.createAt.Before(fileDate.expireAt):
					// rename and append
					fileName := fw.filePathJoin(fw.filePrefix, fileDate.string())
					filePath := fw.filePathJoin(fw.fileDir, fileName)

					if os.Rename(fw.filePath, filePath) == nil {
						list = append(list, &fileInfo{
							fileDate: fileDate,
							fileName: fileName,
							filePath: filePath,
						})
					}
				}
			case strings.HasPrefix(file.Name(), fw.filePrefix):
				fileDate := newFileDateWithStr(strings.TrimPrefix(file.Name(), fw.filePrefix), fw.opts.TTL)
				if fileDate != nil {
					filePath := fw.filePathJoin(fw.fileDir, file.Name())
					switch {
					case current.createAt.After(fileDate.expireAt):
						// remove
						os.Remove(filePath)
					case current.createAt.After(fileDate.createAt) && current.createAt.Before(fileDate.expireAt):
						// append
						list = append(list, &fileInfo{
							fileDate: fileDate,
							fileName: file.Name(),
							filePath: filePath,
						})
					}
				}
			}
		}

		sort.Slice(list, func(i, j int) bool {
			return list[i].fileDate.createAt.Before(list[j].fileDate.createAt)
		})

		for _, fileInfo := range list {
			fw.fileStack.push(fileInfo)
		}
	}

	fw.openCurrentFile(current)

	go fw.daemon()
	return nil
}

// daemon .
func (fw *fileWriter) daemon() {
	next := parseZero(time.Now()).Add(time.Hour * 24)
	timer := time.NewTimer(next.Sub(time.Now()))

	for {
		select {
		case <-timer.C:
			fw.lock.Lock()
			current := newFileDate(fw.opts.TTL)
			fw.openCurrentFile(current)
			fw.lock.Unlock()

			next = next.Add(time.Hour * 24)
			timer = time.NewTimer(next.Sub(time.Now()))
		case <-fw.closing:
			fw.closeCurrentFile()
			return
		}
	}
}

// currentFile .
func (fw *fileWriter) openCurrentFile(currentDate *fileDate) error {
	currentFile, err := os.OpenFile(fw.filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	if fw.writer != nil && fw.current != nil {
		newPath := fw.filePathJoin(fw.fileDir, fw.filePrefix, fw.current.fileDate.string())
		os.Rename(fw.current.filePath, newPath)
		fw.current.filePath = newPath
		fw.fileStack.push(fw.current)

		fw.closeCurrentFile()
	}

	fw.writer = currentFile
	fw.current = &fileInfo{
		fileDate: currentDate,
		fileName: fw.fileName,
		filePath: fw.filePath,
	}
	return nil
}

// filePathJoin return suffix[0] + ... + suffix[n]
func (fw *fileWriter) filePathJoin(args ...string) string {
	var n int
	for _, suffix := range args {
		n += len(suffix)
	}

	var builder strings.Builder
	builder.Grow(n)

	for _, suffix := range args {
		builder.WriteString(suffix)
	}
	return builder.String()
}
