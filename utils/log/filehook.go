package log

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
)

const (
	FireStoping = 0
	FireStop    = 1
	FireHole    = 2
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

func getLevel(s string) logrus.Level {
	switch s {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}

func NewFileHook(level logrus.Level, path string, maxSize int64) (hook *FileHook, err error) {
	return CreateFileHook(path, level, maxSize)
}

//写文件钩子
type FileHook struct {
	levels   []logrus.Level
	date     int
	maxSize  int64
	currSize int64
	path     string
	module   string
	link     string
	file     *os.File
	runFlag  int64
	buf      *bufio.Writer
	queue    chan *logNode
}

type logNode struct {
	level   logrus.Level
	now     time.Time
	file    string
	line    int
	content string
}

func (f *FileHook) Levels() []logrus.Level {
	return f.levels
}

func (f *FileHook) Fire(entry *logrus.Entry) (err error) {
	//初始化
	if entry.Level < logrus.PanicLevel || !atomic.CompareAndSwapInt64(&f.runFlag, FireHole, FireHole) {
		return errors.New("level is outside or log is fire stop")
	}

	node := &logNode{}
	_, node.file, node.line, _ = runtime.Caller(5)
	node.file = filepath.Base(node.file)

	node.content = fmt.Sprintf(
		"MODULE=%s|  message=%v \n",
		entry.Data["MODULE"], entry.Message)
	node.level = entry.Level
	node.now = time.Now()
	f.queue <- node
	return err
}

func CreateFileHook(path string, level logrus.Level, size int64) (f *FileHook, err error) {
	f = new(FileHook)
	if atomic.CompareAndSwapInt64(&f.runFlag, FireHole, FireHole) {
		return nil, errors.New("can't be reused")
	}

	if len(path) < 1 || size < 1 || size > 0x100000000 {
		return nil, errors.New("invalid parameter(size or path)")
	}

	if logrus.Level(level) < logrus.PanicLevel || logrus.Level(level) > logrus.DebugLevel {
		return nil, errors.New("invalid parameter(level)")
	}

	levels := []logrus.Level{}
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}
	f.file = nil
	f.levels = levels
	f.date = 0
	f.maxSize = size
	f.currSize = 0
	f.path = path
	f.module = filepath.Base(os.Args[0])
	f.link = f.path + "/" + f.module + ".log"
	if err := f.loadFile(); err != nil {
		if err = f.createFile(); err != nil {
			return nil, err
		}
	}

	f.runFlag = FireHole
	f.queue = make(chan *logNode, 1024)

	go f.run()
	go f.signalHandler()

	return
}

func (f *FileHook) createFile() (err error) {
	err = os.MkdirAll(f.path, os.ModePerm)
	if err != nil {
		return
	}

	at := time.Now()
	year, mon, day := at.Date()
	newPath := f.path + "/" + f.module + "." + at.Format("20060102.150405.000") + ".log"

	var currSize int64 = 0

	info, err := os.Stat(newPath)
	if err == nil {
		currSize = info.Size()
	}

	file, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_SYNC|os.O_APPEND, 0666)

	if err != nil {
		return
	}

	f.file = file
	f.date = year*10000 + int(mon)*100 + day
	f.currSize = currSize
	f.createLink(newPath)
	return
}

func (f *FileHook) run() {
	ticker := time.NewTicker(time.Microsecond * 500)
	f.buf = bufio.NewWriterSize(f.file, 1048576)

	for {
		select {
		case node := <-f.queue:
			if node.level == 0 {
				f.buf.Flush()
				f.runFlag = FireStop
				return
			}
			if f.switchFile(&node.now, func() {
				f.buf.Flush()
			}) {
				f.buf.Reset(f.file)
			}

			f.processBucket(node, f.buf)
		case <-ticker.C:
			if f.buf.Buffered() > 0 {
				f.buf.Flush()
			}
		}
	}
}

func (f *FileHook) processBucket(node *logNode, iobuf *bufio.Writer) {
	var level = logrus.Level(node.level)

	size, _ := iobuf.WriteString(fmt.Sprintf("[%2s][%02d:%02d:%02d.%03d][%s:%d]: ",
		level.String(),
		node.now.Hour(),
		node.now.Minute(),
		node.now.Second(),
		node.now.Nanosecond()/1000000,
		node.file,
		node.line))
	f.currSize += int64(size)
	size, _ = iobuf.WriteString(node.content)
	f.currSize += int64(size)
}

func (f *FileHook) switchFile(i *time.Time, handler func()) bool {
	year, mon, day := i.Date()
	now := year*10000 + int(mon)*100 + day
	if f.currSize < f.maxSize && now <= f.date {
		return false
	}
	handler()
	if f.file != nil {
		f.file.Close()
		f.file = nil
	}
	err := f.createFile()
	if err != nil {
		return false
	}
	return true
}

func (f *FileHook) loadFile() (err error) {
	filename, err := os.Readlink(f.link)
	if err != nil {
		return
	}

	newPath := f.path + "/" + filename
	info, err := os.Stat(newPath)
	if err != nil {
		return
	}
	//O_SYNC:同步方式打开，即不使用缓存，直接写入硬盘
	file, err := os.OpenFile(newPath, os.O_WRONLY|os.O_SYNC|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	f.file = file
	year, mon, day := info.ModTime().Date()
	f.date = year*10000 + int(mon)*100 + day
	f.currSize = info.Size()
	return
}

func (f *FileHook) createLink(filename string) (err error) {
	os.Remove(f.link)
	return os.Symlink(filepath.Base(filename), f.link)
}

func (f *FileHook) unInitialize() error {
	if !atomic.CompareAndSwapInt64(&f.runFlag, FireHole, FireStoping) {
		return errors.New("Cannot be reused")
	}

	exit_notify := &logNode{}
	exit_notify.level = 0
	f.queue <- exit_notify

	for f.runFlag == FireStoping {
		time.Sleep(time.Millisecond * 10)
	}

	f.file.Close()
	f.file = nil
	f.levels = nil
	f.date = 0
	f.maxSize = 0
	f.currSize = 0
	f.path = ""
	f.module = ""
	f.link = ""
	f.runFlag = FireStop
	f.buf = nil
	close(f.queue)
	return nil
}

func (f *FileHook) signalHandler() {
	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case sig := <-sigChan:
			fmt.Printf("received signal is:%v and exit the whole world", sig)
			f.buf.Flush()
			os.Exit(1)
		}
	}
}
