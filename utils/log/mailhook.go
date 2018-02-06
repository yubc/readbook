package log

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

type MailAuthHook struct {
	Module   string
	AppName  string
	Host     string
	Port     int
	Alias    string
	From     string
	To       []string
	Username string
	Password string
}

func NewMailAuthHook(conf *Configs, Module string) (*MailAuthHook, error) {
	password, _ := base64.StdEncoding.DecodeString(conf.EmailPassword)
	return &MailAuthHook{
		Module:   Module,
		AppName:  Module,
		Host:     conf.EmailHost,
		Port:     conf.EmailPort,
		Alias:    conf.EmailAlias,
		From:     conf.EmailFrom,
		To:       strings.Split(conf.EmailTo, ";"),
		Username: conf.EmailFrom,
		Password: string(password)}, nil
}

func (hook *MailAuthHook) Fire(entry *logrus.Entry) error {
	var now = time.Now()
	_, file, line, _ := runtime.Caller(5)
	file = filepath.Base(file)
	var buf = bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("[%s][%02d:%02d:%02d.%03d][%s:%d]: MODULE=%s|  message=%v \n",
		entry.Level.String(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond()/1000000,
		file,
		line, entry.Data["MODULE"], entry.Message))

	message := []byte("From: " + _Config.EmailAlias + "<" + _Config.EmailFrom + ">\r\nSubject: " + hook.Module + " error" + "\r\n" + "\r\n\r\n" + buf.String())

	auth := smtp.PlainAuth("", hook.Username, hook.Password, hook.Host)
	return smtp.SendMail(
		hook.Host+":"+strconv.Itoa(hook.Port),
		auth,
		_Config.EmailFrom,
		hook.To,
		message,
	)
}

// Levels returns the available logging levels.
func (hook *MailAuthHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}
