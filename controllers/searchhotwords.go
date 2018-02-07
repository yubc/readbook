package controllers

import (
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func SearchHotwords(c *gin.Context) {
	data := NewSetData(c)
	v := models.SearchHotwordError{}

	bd, err := utils.HttpGet(models.SearchHotWord)
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
