package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func UpdateOpenIdToken(code string) (tk *Token, err error) {
	var (
		result struct {
			AccessTokenMsg
			Token
		}
	)
	if tk == nil {
		tk = new(Token)
	}

	URL := fmt.Sprintf(OauthRedirectURL, WxAppId, WxSecret, code)
	bts, err := HttpGet(URL)

	err = json.Unmarshal(bts, &result)
	if err != nil {
		return
	}
	if result.ErrCode != ErrCodeOK {
		return nil, errors.New(result.ErrMsg)
	}

	// 由于网络的延时 和 分布式服务器之间的时间可能不是绝对同步, access_token 过期时间留了一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		return nil, errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 20
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*15:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 20
	default:
		return nil, errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
	}

	tk.AccessToken = result.AccessToken
	tk.CreatedAt = time.Now().Unix()
	tk.ExpiresIn = result.ExpiresIn
	if result.RefreshToken != "" {
		tk.RefreshToken = result.RefreshToken
	}
	if result.OpenId != "" {
		tk.OpenId = result.OpenId
	}

	if result.Scope != "" {
		tk.Scope = result.Scope
	}
	return
}
