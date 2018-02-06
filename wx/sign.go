package wx

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"sort"
	"strings"
)

//验证加密
func ConnectSign(timestamp string, nonce string) string {
	str := []string{GetToken(), timestamp, nonce}
	sort.Strings(str)

	s := sha1.New()
	fmt.Fprint(s, strings.Join(str, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

// Sign2 微信支付签名.
//  params: 待签名的参数集合
//  apiKey: api密钥
//  h:      hash.Hash, 如果为 nil 则默认用 md5.New(), 特别注意 h 必须是 initial state.
func Sign2(params map[string]string, apiKey string, h hash.Hash) string {
	if h == nil {
		h = md5.New()
	}

	keys := make([]string, 0, len(params))
	for k := range params {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	bufw := bufio.NewWriterSize(h, 128)
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		bufw.WriteString(k)
		bufw.WriteByte('=')
		bufw.WriteString(v)
		bufw.WriteByte('&')
	}
	bufw.WriteString("key=")
	bufw.WriteString(apiKey)
	bufw.Flush()

	signature := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}

//shal加密
func ShalSign(content string) string {

	s := sha1.New()
	io.WriteString(s, content)
	return fmt.Sprintf("%x", s.Sum(nil))
}
