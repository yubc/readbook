package routers

import "github.com/gin-gonic/gin"

func Core(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods",
		"OPTIONS,GET,HEAD,POST,PUT,DELETE,TRACE,CONNECT")
	c.Header("Access-Control-Allow-Headers",
		"simple headers,Accept,Accept-Language,Content-Language,Content-Type")
}
