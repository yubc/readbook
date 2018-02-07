package controllers

import (
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func StatisticsController(c *gin.Context) {
	data := NewSetData(c)
	v := models.Statistics{}

	bd, err := utils.HttpGet(models.StatisticsURL)
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

func CatsDetail(c *gin.Context) {
	data := NewSetData(c)
	v := models.CatsDetail{}

	bd, err := utils.HttpGet(models.CatsURL)
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
