package models

const (
	Success   = 0
	ErrorArgs = iota

	ErrorWxServer
	ErrorServer = -1
)

var e = map[int]string{
	Success:   "success",
	ErrorArgs: "参数错误",

	ErrorWxServer: "读取微信数据错误",
	ErrorServer:   "网络开小差了，一会儿再来试试吧~",
}

//获取错误信息
func GetErrMsg(data *Data, errCode int) {
	data.Ret = errCode
	if s, ok := e[errCode]; ok {
		data.Msg = s
	} else {
		data.Ret = ErrorServer
		data.Msg = e[ErrorServer]
	}
	return
}
