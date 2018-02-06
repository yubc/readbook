package controllers

import (
	"fmt"
	"net/http"
	"readbook/models"
	"readbook/wx"
	"time"

	"github.com/gin-gonic/gin"
)

func GetWxOpenId(c *gin.Context) {
	if code := c.Query("code"); code == "" {
		c.Redirect(http.StatusFound, redirectURL(models.Login))
	} else {
		token, err := getWxOpendidFromoauth2(code)
		fmt.Println(token, err)
	}
}

//获取回调地址
func redirectURL(pageType int) string {
	return fmt.Sprintf(wx.OauthURL, models.Conf.WxAppId,
		fmt.Sprintf("%s/readbook/oauth?pt=%d&dt=%d", models.Conf.DominName, pageType, time.Now().Unix()))
}
func getWxOpendidFromoauth2(code string) (string, error) {

	token, err := wx.UpdateOpenIdToken(code)

	return token.OpenId, err
}
