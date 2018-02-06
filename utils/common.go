package utils

import (
	"strings"
)

func IsExistString(str string, strList []string) bool {
	l := len(strList)
	for i := 0; i < l; i++ {
		if strings.Contains(strList[i], str) {
			return true
		}
	}
	return false
}
