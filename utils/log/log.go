package log

import (
	"encoding/json"
	"fmt"

	//"os"

	"github.com/Sirupsen/logrus"
)

var _Config = &Configs{}

func LogFiring(conf *Configs, module string) (Log *logrus.Entry, err error) {

	discard, Module := discardOut{}, module
	_Config = conf

	var l = &logrus.Logger{
		Formatter: &logrus.TextFormatter{FullTimestamp: true},
		Out:       discard,
		Level:     getLevel(_Config.LogLevel),
		Hooks:     make(map[logrus.Level][]logrus.Hook),
	}

	//读取配置文件
	for _, v := range _Config.LogHook {
		if v == "file" {
			fileHook, err := NewFileHook(l.Level, _Config.LogPath, _Config.LogMaxSize)
			if err != nil {
				return Log, err
			}
			fmt.Println("MODULE=log| server start|add filehook")
			l.Hooks.Add(fileHook)

		}
		if v == "mail" {
			//接收level error 以上级别日志
			mailHook, err := NewMailAuthHook(_Config, module)
			if err != nil {
				return Log, err
			}
			fmt.Println("MODULE=log| server start|add mailhook")
			l.Hooks.Add(mailHook)
		}
	}

	//l.Hooks.Add(new(LineNumHook))

	var field = logrus.Fields{}
	field["MODULE"] = Module
	Log = logrus.NewEntry(l).WithFields(field)
	return
}

func LogInit(b []byte, module string) {

	err := json.Unmarshal(b, _Config)
	if err != nil {
		return
	}
	fileHook, err := NewFileHook(logrus.DebugLevel, _Config.LogPath, _Config.LogMaxSize)
	if err != nil {
		panic(err)
	}

	fmt.Println("MODULE=log| server start|add filehook")
	logrus.AddHook(fileHook)
	//接收level error 以上级别日志
	mailHook, err := NewMailAuthHook(_Config, module)
	if err != nil {
		panic(err)
	}
	fmt.Println("MODULE=log| server start|add mailhook")
	logrus.AddHook(mailHook)

	logrus.WithField("MODULE", module)
}

type Configs struct {
	LogPath       string   //日志文件目录
	LogLevel      string   //日志级别
	LogMaxSize    int64    //单个日志文件容量
	LogBuffSize   int      //日志缓存容量
	LogHook       []string //日志钩子
	EmailHost     string
	EmailPort     int
	EmailFrom     string
	EmailTo       string
	EmailPassword string
	EmailAlias    string //发件人别名
}
