package main

import (
	"fmt"
	_ "readbook/conf"
	"readbook/inits"
	"readbook/models"
	"readbook/routers"
	"readbook/wx"

	"github.com/google/gops/agent"

	"github.com/gin-gonic/gin"
)

var (
	engine  *gin.Engine
	Version string
)

func main() {
	fmt.Println("start version:", Version)
	engine = gin.Default()

	inits.LoadInit()
	wx.InitKey(models.Conf.WxAppId, models.Conf.WxSecret,
		models.Conf.WxPayMchid, models.Conf.WxPayKey, models.Conf.WxToken)

	go func() {
		err := agent.Listen(agent.Options{})
		if err != nil {
			fmt.Println(err)
		}
	}()
	routers.Router(engine)
	engine.Run(models.Conf.RunPort)
}
