package wx

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func UnifiedOrder2(req *UnifiedOrderRequest) (resp *UnifiedOrderResponse, err error) {
	m1 := make(map[string]string, 24)
	m1["body"] = req.Body
	m1["out_trade_no"] = req.OutTradeNo
	m1["total_fee"] = strconv.FormatInt(req.TotalFee, 10)
	m1["spbill_create_ip"] = req.SpbillCreateIP
	m1["notify_url"] = req.NotifyURL
	m1["trade_type"] = req.TradeType
	if req.DeviceInfo != "" {
		m1["device_info"] = req.DeviceInfo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = GetRandString(10)
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}
	if req.Detail != "" {
		m1["detail"] = req.Detail
	}
	if req.Attach != "" {
		m1["attach"] = req.Attach
	}
	if req.FeeType != "" {
		m1["fee_type"] = req.FeeType
	}

	if req.GoodsTag != "" {
		m1["goods_tag"] = req.GoodsTag
	}
	if req.ProductId != "" {
		m1["product_id"] = req.ProductId
	}
	if req.LimitPay != "" {
		m1["limit_pay"] = req.LimitPay
	}
	if req.OpenId != "" {
		m1["openid"] = req.OpenId
	}
	if req.SubOpenId != "" {
		m1["sub_openid"] = req.SubOpenId
	}
	if req.SceneInfo != "" {
		m1["scene_info"] = req.SceneInfo
	}

	m2, err := UnifiedOrder(m1)
	if err != nil {
		return nil, err
	}

	// 校验 trade_type
	respTradeType := m2["trade_type"]
	if respTradeType != req.TradeType {
		err = fmt.Errorf("trade_type mismatch, have: %s, want: %s", respTradeType, req.TradeType)
		return nil, err
	}

	resp = &UnifiedOrderResponse{
		PrepayId:   m2["prepay_id"],
		TradeType:  respTradeType,
		DeviceInfo: m2["device_info"],
		CodeURL:    m2["code_url"],
		MWebURL:    m2["mweb_url"],
		AppId:      m2["appid"],
		NonceStr:   m2["nonce_str"],
		PaySign:    m2["sign"],
	}
	return resp, nil
}

// UnifiedOrder 统一下单.
func UnifiedOrder(req map[string]string) (resp map[string]string, err error) {

	if strings.TrimSpace(WxPayAppId) == "" || strings.TrimSpace(WxPayMchId) == "" {
		err = fmt.Errorf("%s", "invalid appid or mchid")
		return
	}
	if req["appid"] == "" {
		req["appid"] = WxPayAppId
	}
	if req["mch_id"] == "" {
		req["mch_id"] = WxPayMchId
	}

	// 获取请求参数的 sign_type 并检查其有效性
	var reqSignType string
	switch signType := req["sign_type"]; signType {
	case "", SignType_MD5:
		reqSignType = SignType_MD5
	case SignType_HMAC_SHA256:
		reqSignType = SignType_HMAC_SHA256
	default:
		return nil, fmt.Errorf("unsupported request sign_type: %s", signType)
	}

	// 如果没有签名参数补全签名
	if req["sign"] == "" {
		switch reqSignType {
		case SignType_MD5:
			req["sign"] = Sign2(req, WxPayKey, md5.New())
		case SignType_HMAC_SHA256:
			req["sign"] = Sign2(req, WxPayKey, hmac.New(sha256.New, []byte(WxPayAppId)))
		}
	}

	buffer := new(bytes.Buffer)
	if err = EncodeXMLFromMap(buffer, req, "xml"); err != nil {
		return nil, err
	}
	body := buffer.Bytes()

	log.Printf("支付:[订单时间:%v] [订单消息:%s]\n", time.Now().Format("2006-01-02 15:04:05"), string(body))
	resp, err = postXML(WxPayUnifyOrderUrl, body, reqSignType)
	if err != nil {
		return nil, err
	}
	log.Printf("支付:[订单时间:%v] [订单消息:%s]\n", time.Now().Format("2006-01-02 15:04:05"), resp)
	return resp, nil
}

func postXML(url string, body []byte, reqSignType string) (resp map[string]string, err error) {

	httpResp, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	resp, err = DecodeXMLHttpResponse(httpResp.Body)
	if err != nil {
		return nil, err
	}

	// 判断协议状态
	returnCode := resp["return_code"]
	if returnCode == "" {
		return nil, ErrNotFoundReturnCode
	}
	if returnCode != ReturnCodeSuccess {
		return nil, &Error{
			ReturnCode: returnCode,
			ReturnMsg:  resp["return_msg"],
		}
	}
	return resp, nil
}
