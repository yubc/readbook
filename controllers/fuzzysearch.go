package controllers

import (
	"fmt"
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func FuzzySearch(c *gin.Context) {
	data := NewSetData(c)
	query, ok := c.GetQuery("query")
	if !ok {
		data.Ret = models.ErrorArgs
		return
	}
	v := models.FuzzySearch{}
	bd, err := utils.HttpGet(fmt.Sprintf(models.FuzzySearchURL, utils.URLQueryEscape(query)))
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
