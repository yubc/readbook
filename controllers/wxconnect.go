package controllers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"readbook/models"
	"readbook/utils"
	"readbook/wx"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 验证
func ConnectWx(c *gin.Context) {

	timestamp, nonce, signatureIn := c.Query("timestamp"), c.Query("nonce"), c.Query(("signature"))
	//验证
	signatureGen := wx.ConnectSign(timestamp, nonce)
	if signatureGen != signatureIn {
		c.String(http.StatusOK, "")
		return
	}
	echostr := c.Query("echostr")
	c.String(http.StatusOK, echostr)
}

//菜单点击事件
func WxEvent(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		models.Log.Infoln("微信菜单用户点击、扫码解析body错误", err)
		return
	}
	wxEvent := wx.WxMenuEvent{}
	if err := xml.Unmarshal(body, &wxEvent); err == nil {
		//处理菜单点击
		fmt.Println(">>>>>", fmt.Sprintf("%+v", wxEvent))
		dealWithMenuEvent(&wxEvent)
	} else {
		models.Log.Infoln("微信菜单用户点击、扫码响应错误", err)
	}
}

//获取token
func GetAccessToken(c *gin.Context) {
	data := NewSetData(c)

	ss := sessions.Default(c)

	v := ss.Get("token").(string)
	if v != "" {
		data.Data["data"] = v
		return
	}
	token, err := getToken()
	if err != nil {
		ss.Set("token", token)
	}
	ss.Save()
	data.Data["data"] = token
}

//创建菜单
func CreatMenu(c *gin.Context) {

	wxData := wx.AccessTokenMsg{}
	data := NewSetData(c)

	menuStr := `{
    "button": [
    {
        "name": "进入阅读",
        "type": "view",
        "url": "http://test.renrenmeishu.com/static/index.html"
    }
    ]
  }`
	token := c.GetString("token")
	if token == "" {
		ss := sessions.Default(c)
		tokenInterface := ss.Get("token")
		if tokenInterface == nil {
			tokens, err := getToken()
			if err == nil {
				token = tokens
				ss.Set("token", tokens)
				ss.Save()
			} else {
				return
			}
		} else {
			token = tokenInterface.(string)
		}
	}

	body, err := utils.HttpPost(fmt.Sprintf(wx.MenuURL, token), []byte(menuStr), true)
	if err != nil {
		data.Ret = models.ErrorWxServer
		return
	}
	json.Unmarshal(body, &wxData)
	models.Log.Infof("[创建微信菜单] %+v", wxData)
	data.Data["data"] = "创建成功!"
}

//定时读取
func StartTimer() {

	tnow := time.NewTicker(1 * time.Minute)
	for v := range tnow.C {
		if v.Minute() == 0 && v.Second() == 0 {
			getToken()
		}
	}
}

//获取基本token
func getToken() (string, error) {

	var result struct {
		wx.AccessToken
		wx.AccessTokenMsg
	}

	body, err := utils.HttpGet(fmt.Sprintf(wx.AccessTokenURL, models.Conf.WxAppId, models.Conf.WxSecret))
	if err != nil {
		return "", err
	}

	models.Log.Infof("[GetAccessToken] [%+v]", string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		models.Log.Infof("[GetAccessToken] [%v]", err)
	}

	if result.ErrCode != wx.ErrCodeOK {
		return "", errors.New(result.ErrMsg)
	}

	// 由于网络的延时 和 分布式服务器之间的时间可能不是绝对同步, access_token 过期时间留了一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		return "", errors.New("expires_in too large: " + strconv.FormatInt(int64(result.ExpiresIn), 10))
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 20
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*15:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 20
	default:
		return "", errors.New("expires_in too small: " + strconv.FormatInt(int64(result.ExpiresIn), 10))
	}

	return result.Token, nil
}

func dealWithMenuEvent(wxEvent *wx.WxMenuEvent) {
	switch wxEvent.Event {
	case "VIEW": //点击菜单底部事件
	case "CLICK": //点击按钮
		if strings.EqualFold(wxEvent.Event, "") {
			//点击用户中心
		} else if strings.EqualFold(wxEvent.Event, "") {
			//点击公告
		}
	case "subscribe": //扫码
		if strings.Contains(wxEvent.EventKey, "qrscene") {
			//未关注，扫描其他人首次关注
			upperid := strings.Split(wxEvent.EventKey, "_")[1]
			fmt.Println(">>>>>upperid", upperid)
		} else {
			//关注 扫描官方二维码
		}
	case "SCAN": //已关注，扫描他人二维码

	}
}
