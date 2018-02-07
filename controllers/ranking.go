package controllers

import (
	"fmt"
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func Ranking(c *gin.Context) {
	data := NewSetData(c)
	v := models.RankingInfo{}

	rankid := c.Param("rankid")
	if rankid == "" {
		data.Ret = models.ErrorArgs
		return
	}

	bd, err := utils.HttpGet(fmt.Sprintf(models.RankIdURL, rankid))
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
