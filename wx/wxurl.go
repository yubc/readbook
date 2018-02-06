package wx

const (
	WxPayUnifyOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder" //下订单
	WxPayQueryOrderUrl = "https://api.mch.weixin.qq.com/pay/orderquery"   //订单查询
)

var (
	AccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	MenuURL        = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"

	OauthURL         = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
	OauthRedirectURL = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"

	TemplatePushURL = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
	MsgPushURL      = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"

	QrcodeTicket  = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	QrcodeShowURL = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"

	JssdkURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)
