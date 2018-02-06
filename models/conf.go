package models

type Configs struct {
	RunPort        string     `yaml:"runport"`        //运行端口
	DominName      string     `yaml:"dominname"`      //服务地址
	EnvMode        string     `yaml:"envmode"`        //运行模式 dev prod
	ProjectPath    string     `yaml:"projectpath"`    //运行路径
	WxAppId        string     `yaml:"wxappid"`        //微信公众号appid
	WxSecret       string     `yaml:"wxsecret"`       //微信公众号appsecret
	WxPayMchid     string     `yaml:"wxpaymchid"`     //微信公众号付款machid
	WxPayKey       string     `yaml:"wxpaykey"`       //微信公众号appid
	WxPushTemplate []string   `yaml:"wxpushtemplate"` //微信公众号appid
	WxToken        string     `yaml:"wxtoken"`        //微信公众号token
	WxTemplate     []Template `yaml:"-"`

	LogProject    string   `yaml:"logproject"`  //项目名
	LogPath       string   `yaml:"logpath"`     //日志文件目录
	LogLevel      string   `yaml:"loglevel"`    //日志级别
	LogMaxSize    int64    `yaml:"logmaxsize"`  //单个日志文件容量
	LogBuffSize   int      `yaml:"logbuffsize"` //日志缓存容量
	LogHook       []string `yaml:"loghook"`     //日志钩子
	EmailHost     string   `yaml:"emailhost"`
	EmailPort     int      `yaml:"emailport"`
	EmailFrom     string   `yaml:"emailfrom"`
	EmailTo       string   `yaml:"emailto"`
	EmailPassword string   `yaml:"emailpassword"`
	EmailAlias    string   `yaml:"emailalias"` //发件人别名
}

//推送数据
type Template struct {
	KeyType string
	KeyId   string
	KeyURL  string
}
