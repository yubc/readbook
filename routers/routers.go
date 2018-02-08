package routers

import (
	"io/ioutil"
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
	router.Use(Core)

	book := router.Group("/readbook")
	{
		book.POST("/wx", ctl.WxEvent)
		book.GET("/wx", ctl.ConnectWx)

		book.GET("/oauth", ctl.GetWxOpenId)
		book.GET("/createmenu", ctl.CreatMenu)

		book.GET("gettoken", ctl.GetAccessToken)

		book.GET("/cats/lv2/statistics", ctl.StatisticsController) //获取所有分类
		book.GET("/catsdetail/lv2", ctl.CatsDetail)                //细分分类

		book.GET("/ranking/gender", ctl.RankGenderController) //获取排行榜类型

		book.GET("/book/search-hotwords", ctl.SearchHotwords) //搜索热词

		book.GET("book/auto-complete", ctl.AutoComplete) //自动查找
		book.GET("book/fuzzy-search", ctl.FuzzySearch)   //模糊查找

		book.GET("/rankinglist/:rankid", ctl.Ranking)   //获取排行榜小说
		book.GET("/book/by-categories", ctl.Categories) //根据分类获取小说列表

		book.GET("/bookinfo/:bookId", ctl.BookInfo) //获取小说信息
		book.GET("/btoc", ctl.BtocSummary)          //获取小说正版源
		book.GET("/atoc", ctl.AtocSummary)          //获取小说正版源于盗版源(混合)

		book.GET("/atocchapter/:sourceId", ctl.GetChapter) //获取小说章节(根据小说源id)
		book.GET("/getChapter", ctl.GetChapterContent)     //获取小说内容

		book.GET("/bookdetail", ctl.NewBookChapter) //获取最新章节

		book.GET("/index", func(c *gin.Context) {
			c.Header("Content-Type", "text/html; charset=utf-8")
			v, _ := ioutil.ReadFile(path.Join(models.Conf.ProjectPath, "/static/index.html"))
			c.String(200, string(v))
		})
	}
	router.Static("/static", path.Join(models.Conf.ProjectPath, "static"))

}
