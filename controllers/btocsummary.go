package controllers

import (
	"fmt"
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func BtocSummary(c *gin.Context) {
	data := NewSetData(c)
	view, book := c.Query("view"), c.Query("book")
	if view == "" || book == "" {
		data.Ret = models.ErrorArgs
		return
	}
	v := []models.SummaryList{}

	bd, err := utils.HttpGet(fmt.Sprintf(models.BtocURL, view, book))
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

func AtocSummary(c *gin.Context) {
	data := NewSetData(c)
	view, book := c.Query("view"), c.Query("book")
	if view == "" || book == "" {
		data.Ret = models.ErrorArgs
		return
	}
	v := []models.SummaryList{}

	bd, err := utils.HttpGet(fmt.Sprintf(models.AtocURL, view, book))
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
