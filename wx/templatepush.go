package wx

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

/***
模版推送
*/

func WxTemplatePush(wxTemplate WxTemplateData) (wr WxMsgReturn) {
	wr = WxMsgReturn{
		Errcode: 0,
		Errmsg:  "",
	}
	if len(strings.TrimSpace(wxTemplate.AccessToken)) == 0 {
		wr.Errcode = -1
		wr.Errmsg = ErrNotInvalidToken.Error()
		return
	}

	sendMsg, err := json.Marshal(wxTemplate)
	if err != nil {
		wr.Errcode = -1
		wr.Errmsg = err.Error()
		return
	}
	bts, err := HttpPost(fmt.Sprintf(TemplatePushURL, wxTemplate.AccessToken), sendMsg)
	if err != nil {
		wr.Errcode = -1
		wr.Errmsg = err.Error()
		return
	}
	json.Unmarshal(bts, &wr)
	log.Printf("[push:%v data:%+v]", time.Now(), wr)
	return
}
