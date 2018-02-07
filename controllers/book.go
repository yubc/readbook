package controllers

import (
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func NewBookChapter(c *gin.Context) {
	data := NewSetData(c)
	id, ok := c.GetQuery("id")
	if !ok {
		data.Ret = models.ErrorArgs
		return
	}
	v := models.Gender{}

	bd, err := utils.HttpGet(models.BookLastChapter + id)
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
