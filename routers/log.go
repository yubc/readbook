package routers

import (
	"io/ioutil"
	"time"

	"readbook/models"

	"github.com/gin-gonic/gin"
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

func defaultLog(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path

	c.Next()

	latency := time.Since(start)
	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()
	statusColor := colorForStatus(statusCode)
	methodColor := colorForMethod(method)
	var args string

	if c.Request.Header.Get("Content-Type") == "application/json" {
		b, _ := ioutil.ReadAll(c.Request.Body)
		args = string(b)
	} else {
		c.Request.ParseForm()
		args = c.Request.Form.Encode()
	}

	models.Log.Infof("|%s %3d %s| %6v | %s |%s %s %s %s %s",
		statusColor, statusCode, reset,
		latency,
		clientIP,
		methodColor, method, reset,
		path,
		args,
	)
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}
