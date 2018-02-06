package wx

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
获取微信二维码
*/

func getQrFromWx(openid string, token string) (string, error) {

	ticker := ""

	if strings.TrimSpace(token) == "" || strings.TrimSpace(openid) == "" {
		return "", fmt.Errorf("invalid %s %s", openid, token)
	}
	qrGet := QrGet{
		ActionName: "QR_LIMIT_SCENE",
		AInfo: &QrActionInfo{
			Sc: &Scene{SceneStr: openid},
		},
	}
	jsonBytes, err := json.Marshal(qrGet)
	if err != nil {
		return ticker, err
	}
	URL := fmt.Sprintf(QrcodeTicket, token)
	bts, err := HttpPost(URL, jsonBytes)
	if err != nil {
		return ticker, err
	}

	qrBody := new(QrBody)
	if err := json.Unmarshal(bts, qrBody); err == nil {
		fmt.Printf("qrbody-----%+v\n", qrBody)
		ticker = qrBody.Ticket
	} else {
		fmt.Println("json转换错误", err)
		return ticker, err
	}
	return ticker, nil
}
