package controllers

import (
	"fmt"
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func BookInfo(c *gin.Context) {
	data := NewSetData(c)
	v := models.BookInfos{}

	bookid := c.Param("bookId")
	if bookid == "" {
		data.Ret = models.ErrorArgs
		return
	}

	bd, err := utils.HttpGet(fmt.Sprintf(models.BookInfoURL, utils.URLQueryEscape(bookid)))
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
