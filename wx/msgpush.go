package wx

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/***
客服消息推送*/

func SendMsg(toUserOpenid string, msg string, token string) {

	msgSend := WxMsg{
		Touser:  toUserOpenid,
		Msgtype: "text",
		Text:    WxMsgText{Content: msg},
	}
	body, err := json.Marshal(msgSend)
	if err != nil {
		return
	}
	msgPost(token, body)
}

func msgPost(token string, jsonbody []byte) error {

	if strings.TrimSpace(token) == "" {
		return ErrNotInvalidToken
	}
	URL := fmt.Sprintf(MsgPushURL, token)

	bts, err := HttpPost(URL, jsonbody)

	if err != nil {
		log.Printf("%s 请求失败 %v\n", "客服向用户发送信息", err)
		return err
	}
	log.Printf("%s 解析结果 %v", "客服向用户发送信息", string(bts))

	return nil
}
