package controllers

import (
	"readbook/models"

	"github.com/gin-gonic/gin"
)

func newData() *models.Data {
	data := &models.Data{
		Ret:  models.Success,
		Data: make(map[string]interface{}),
	}
	return data
}

func NewSetData(c *gin.Context) (data *models.Data) {

	data = newData()
	data.Data["data"] = ""
	c.Set("data", data)
	return
}

//获取数据
func GetData(c *gin.Context) *models.Data {
	d, exist := c.Get("data")
	if !exist {
		models.Log.Errorf("%s |func=%s |message=%v",
			"get data from http.context", "GetData", "data not exists")

		data := newData()
		models.GetErrMsg(data, models.ErrorServer)
		return data
	}
	dd, ok := d.(*models.Data)
	if !ok {
		models.Log.Errorf("%s |func=%s |message=%v",
			"get data from http.context", "GetData", "data not exists")
	}
	dd.ReturnData = dd.Data["data"]

	return dd
}
