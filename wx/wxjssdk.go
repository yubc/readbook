package wx

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/**
jssdk获取*/
func getJssdk(token string) (Jssdk, error) {

	jssdk := Jssdk{}

	URL := fmt.Sprintf(JssdkURL, token)
	bts, err := HttpGet(URL)
	if err != nil {
		return jssdk, fmt.Errorf("jssdk get err %v", err)
	}
	err = json.Unmarshal(bts, &jssdk)
	return jssdk, err
}

//jssdk签名
func GetJssdkSign(token string, signURL string) (JssdkBack, error) {
	data := JssdkBack{}
	if strings.TrimSpace(token) == "" {
		return data, ErrNotInvalidToken
	}

	v, err := getJssdk(token)
	if err != nil {
		return data, err
	}
	data = JssdkBack{
		Appid:     WxAppId,
		Secret:    WxSecret,
		Timestamp: time.Now().Unix(),
		Noncestr:  GetRandString(10),
	}
	sigstring := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s",
		v.Ticket, data.Noncestr, data.Timestamp, signURL)
	data.Signature = ShalSign(sigstring)

	return data, nil
}
