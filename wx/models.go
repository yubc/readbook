package wx

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

const (
	ErrCodeOK = 0

	ReturnCodeSuccess = "SUCCESS"
	ReturnCodeFail    = "FAIL"

	ResultCodeSuccess = "SUCCESS"
	ResultCodeFail    = "FAIL"

	SignType_MD5         = "MD5"
	SignType_SHA1        = "SHA1"
	SignType_HMAC_SHA256 = "HMAC-SHA256"
)

var (
	WxAppId    = "" //微信公众号appid
	WxSecret   = "" //微信secret
	WxPayAppId = "" //微信支付appid
	WxPayMchId = "" //商家macid
	WxPayKey   = "" //商家apikey
	WxToken    = "" //token

	ErrNotFoundReturnCode = errors.New("not found return_code parameter")
	ErrNotInvalidToken    = errors.New("weixin invalid token")
)

type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

type AccessTokenMsg struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//点击事件
type WxMenuEvent struct {
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   int64         `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`
	Event        string        `xml:"Event"` //VIEW
	EventKey     string        `xml:"EventKey"`
	MenuId       string        `xml:"MenuId"`
	ScanCodeInfo *ScanCodeInfo `xml:"ScanCodeInfo"` //专属于扫码
	Content      string        `xml:"Content"`
	Ticket       string        `xml:"Ticket"`
}
type ScanCodeInfo struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}

//网页授权
type Token struct {
	AccessToken  string `json:"access_token"`            // 网页授权接口调用凭证
	CreatedAt    int64  `json:"created_at"`              // access_token 创建时间, unixtime, 分布式系统要求时间同步, 建议使用 NTP
	ExpiresIn    int64  `json:"expires_in"`              // access_token 接口调用凭证超时时间, 单位: 秒
	RefreshToken string `json:"refresh_token,omitempty"` // 刷新 access_token 的凭证

	OpenId string `json:"openid,omitempty"`
	Scope  string `json:"scope,omitempty"` // 用户授权的作用域, 使用逗号(,)分隔
}

type Error struct {
	XMLName    struct{} `xml:"xml"                  json:"-"`
	ReturnCode string   `xml:"return_code"          json:"return_code"`
	ReturnMsg  string   `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
}

func (e *Error) Error() string {
	bs, err := xml.Marshal(e)
	if err != nil {
		return fmt.Sprintf("return_code: %q, return_msg: %q", e.ReturnCode, e.ReturnMsg)
	}
	return string(bs)
}

//支付
type (
	UnifiedOrderRequest struct {
		XMLName struct{} `xml:"xml" json:"-"`

		// 必选参数
		Body           string `xml:"body"`             // 商品或支付单简要描述
		OutTradeNo     string `xml:"out_trade_no"`     // 商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
		TotalFee       int64  `xml:"total_fee"`        // 订单总金额，单位为分，详见支付金额
		SpbillCreateIP string `xml:"spbill_create_ip"` // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
		NotifyURL      string `xml:"notify_url"`       // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
		TradeType      string `xml:"trade_type"`       // 取值如下：JSAPI，NATIVE，APP，详细说明见参数规定

		// 可选参数
		DeviceInfo string    `xml:"device_info"` // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
		NonceStr   string    `xml:"nonce_str"`   // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
		SignType   string    `xml:"sign_type"`   // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
		Detail     string    `xml:"detail"`      // 商品名称明细列表
		Attach     string    `xml:"attach"`      // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
		FeeType    string    `xml:"fee_type"`    // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
		TimeStart  time.Time `xml:"time_start"`  // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
		TimeExpire time.Time `xml:"time_expire"` // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则
		GoodsTag   string    `xml:"goods_tag"`   // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
		ProductId  string    `xml:"product_id"`  // trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
		LimitPay   string    `xml:"limit_pay"`   // no_credit--指定不能使用信用卡支付
		OpenId     string    `xml:"openid"`      // rade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。
		SubOpenId  string    `xml:"sub_openid"`  // trade_type=JSAPI，此参数必传，用户在子商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid。
		SceneInfo  string    `xml:"scene_info"`  // 该字段用于上报支付的场景信息,针对H5支付有以下三种场景,请根据对应场景上报,H5支付不建议在APP端使用，针对场景1，2请接入APP支付，不然可能会出现兼容性问题
	}

	UnifiedOrderResponse struct {
		XMLName struct{} `xml:"xml" json:"-"`

		// 必选返回
		PrepayId  string `xml:"prepay_id"`  // 微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
		TradeType string `xml:"trade_type"` // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，详细说明见参数规定

		AppId    string `xml:"appid"`
		NonceStr string `xml:"nonce_str"`
		PaySign  string `xml:"paySign"`
		// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
		DeviceInfo string `xml:"device_info"` // 调用接口提交的终端设备号。
		CodeURL    string `xml:"code_url"`    // trade_type 为 NATIVE 时有返回，可将该参数值生成二维码展示出来进行扫码支付
		MWebURL    string `xml:"mweb_url"`    // trade_type 为 MWEB 时有返回
	}

	//支付回调
	NotifyResp struct {
		ReturnCode    string `xml:"return_code"`
		ReturnMsg     string `xml:"return_msg"`
		AppId         string `xml:"appid"`
		MchId         string `xml:"mch_id"`
		Nonce         string `xml:"nonce_str"`
		Sign          string `xml:"sign"`
		ResultCode    string `xml:"result_code"`
		Openid        string `xml:"openid"`
		IsSubscribe   string `xml:"is_subscribe"`
		TradeType     string `xml:"trade_type"`
		BankType      string `xml:"bank_type"`
		FeeType       string `xml:"fee_type"`
		CashFeeType   string `xml:"cash_fee_type"`
		TransactionId string `xml:"transaction_id"`
		OutTradeNo    string `xml:"out_trade_no"`
		Attach        string `xml:"attach"`
		TimeEnd       string `xml:"time_end"`
		TotalFee      int    `xml:"total_fee"`
		CashFee       int    `xml:"cash_fee"`
	}
)

//消息模版推送
type (
	WxTemplateData struct {
		Touser     string                     `json:"touser"`      //用户openid
		TemplateId string                     `json:"template_id"` //模版id
		Url        string                     `json:"url"`
		TopColor   string                     `json:"top_color,omitempty"`
		Data       map[string]*WxTemplateInfo `json:"data"` //string就是模版的key

		AccessToken string `json:"-"`
	}

	WxTemplateInfo struct {
		Value string `json:"value"`
		Color string `json:"color,omitempty"`
	}

	WxMsgReturn struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		Msgid   int    `json:"msgid"`
	}
)

//客服消息
type (
	WxMsg struct {
		Touser  string    `json:"touser"`
		Msgtype string    `json:"msgtype"`
		Text    WxMsgText `json:"text"`
	}
	WxMsgText struct {
		Content string `json:"content"`
	}
)

//二维码
type (
	QrGet struct {
		ActionName string        `json:"action_name"`
		AInfo      *QrActionInfo `json:"action_info"`
	}
	QrActionInfo struct {
		Sc *Scene `json:"scene"`
	}

	Scene struct {
		SceneStr string `json:"scene_str"`
	}

	QrBody struct {
		Ticket        string `json:"ticket"`
		ExpireSeconds int    `json:"expire_seconds"`
		URL           string `json:"url"`
	}
)

//JSSDK
type (
	Jssdk struct {
		Ticket   string `json:"ticket"`
		ExpireIn int    `json:"expires_in"`
		ErrCode  int    `json:"errcode"`
		ErrMsg   string `json:"errmsg"`
	}
	JssdkBack struct {
		Appid     string `json:"appid"`
		Secret    string `json:"secret"`
		Timestamp int64  `json:"timestamp,string"`
		Noncestr  string `json:"noncestr"`
		Signature string `json:"signature"`
	}
)

//初始化key
func InitKey(appid string, secret string, macid string, key string, token string) {
	WxAppId = appid
	WxSecret = secret
	WxPayAppId = appid
	WxPayMchId = macid
	WxPayKey = key
	WxToken = token
}
