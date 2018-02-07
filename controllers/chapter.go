package controllers

import (
	"fmt"
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func GetChapter(c *gin.Context) {
	data := NewSetData(c)
	view, book := c.Param("sourceId"), c.Query("view")
	if view == "" || book == "" {
		data.Ret = models.ErrorArgs
		return
	}
	v := models.ChapterInfo{}

	bd, err := utils.HttpGet(fmt.Sprintf(models.SourceURL, view, book))
	if err != nil {
		data.Ret = models.ErrorNoData
		return
	}
	err = utils.Json().Unmarshal(bd, &v)
	if err != nil {
		data.Ret = models.ErrorJsonInvilid
		return
	}

	data.Data["data"] = v
}

func GetChapterContent(c *gin.Context) {
	data := NewSetData(c)
	chapter, ok := c.GetQuery("chapterUrl")
	if !ok {
		data.Ret = models.ErrorArgs
		return
	}
	v := models.ChapterContentInfo{}

	bd, err := utils.HttpGet(fmt.Sprintf(models.ContentURL, chapter))
	if err != nil {
		data.Ret = models.ErrorNoData
		return
	}
	err = utils.Json().Unmarshal(bd, &v)
	if err != nil {
		data.Ret = models.ErrorJsonInvilid
		return
	}

	data.Data["data"] = v
}
