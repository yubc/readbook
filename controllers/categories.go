package controllers

import (
	"fmt"
	"readbook/models"
	"readbook/utils"

	"github.com/gin-gonic/gin"
)

func Categories(c *gin.Context) {
	data := NewSetData(c)
	v := models.CategoriesList{}

	gender, types, major, minor, start, limit := c.Query("gender"), c.Query("type"), c.Query("major"), c.Query("minor"), c.Query("start"), c.Query("limit")
	if gender == "" || types == "" || major == "" {
		data.Ret = models.ErrorArgs
		return
	}
	if start == "" || limit == "" {
		start = "0"
		limit = "10"
	}
	bd, err := utils.HttpGet(fmt.Sprintf(models.CategoriesURL,
		utils.URLQueryEscape(gender), utils.URLQueryEscape(types),
		utils.URLQueryEscape(major), utils.URLQueryEscape(minor), start, limit))
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
