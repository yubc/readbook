package routers

import (
	"path"
	ctl "readbook/controllers"
	"readbook/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const _SESSION_STORE_KEY = "readbook"

func Router(router *gin.Engine) {

	store := sessions.NewCookieStore([]byte(_SESSION_STORE_KEY))

	router.Use(sessions.Sessions("readingbook", store))
	router.Use(defaultLog)
	router.Use(CommonReturn)

	book := router.Group("/readbook")
	{
		book.POST("/wx", ctl.WxEvent)
		book.GET("/wx", ctl.ConnectWx)

		book.GET("/oauth", ctl.GetWxOpenId)
		book.GET("/createmenu", ctl.CreatMenu)

		book.GET("gettoken", ctl.GetAccessToken)
	}
	router.Static("/static", path.Join(models.Conf.ProjectPath, "static"))
}
