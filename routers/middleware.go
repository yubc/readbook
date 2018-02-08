package routers

import (
	"net/http"
	ctl "readbook/controllers"
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

var Fileter = []string{
	"/readbook/wx",
	"/readbook/oauth",
	"/static",
	"/index",
}

func CommonReturn(c *gin.Context) {
	c.Next()

	Uri := c.Request.URL.Path
	if ok := utils.IsExistString(Uri, Fileter); ok {
		return
	}
	if c.Writer.Status() == 404 {
		c.Abort()
		return
	}

	data := ctl.GetData(c)
	if len(data.Msg) == 0 {
		models.GetErrMsg(data, data.Ret)
	}

	//统一返回数据
	c.JSON(http.StatusOK, data)
}
