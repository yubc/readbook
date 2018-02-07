package utils

import (
	"net/url"
	"strings"

	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func IsExistString(str string, strList []string) bool {
	l := len(strList)
	for i := 0; i < l; i++ {
		if strings.Contains(strList[i], str) {
			return true
		}
	}
	return false
}

func Json() jsoniter.API {
	return json
}

func URLQueryEscape(str string) string {
	return url.QueryEscape(str)
}
