package models

import (
	"github.com/Sirupsen/logrus"
)

const (
	LONGTIMEFORMAT  = "20060102150405"
	STDTIMEFORMAT   = "2006-01-02 15:04:05"
	SHORTTIMEFORMAT = "2006-01-02"
	FMTTIMEFORMAT   = "20060102"
	EXPIRETIME      = 24 * 3600
	ContentType     = "application/json;charset=utf-8"
)
const (
	_ = iota
	Login

	LoginHtml = ""
)

var (
	Conf Configs
	Log  *logrus.Entry
)

type Data struct {
	Ret        int                    `json:"ret,string"`
	Msg        string                 `json:"msg"`
	Data       map[string]interface{} `json:"-"`
	ReturnData interface{}            `json:"data,omitempty"`
}
