package wx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	r       = rand.New(rand.NewSource(time.Now().UnixNano()))
)

//随机数
func GetRandString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	now := time.Now()
	r = rand.New(rand.NewSource(now.UnixNano()))
	return strings.ToUpper(fmt.Sprintf("%x%02x",
		uint64((now.Year()-2006)*100000+int(now.Month())*10000+now.Day()*1000+now.Hour()*100+now.Minute()*10+now.Second()),
		uint64(r.Int31n(0xFFFFF))) + string(b))
}

func HttpGet(URL string) (body []byte, err error) {

	var (
		resp *http.Response
	)
	resp, err = http.Get(URL)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("StatusCode is %d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpPost(URL string, b []byte) (body []byte, err error) {

	var (
		resp *http.Response
		req  *http.Request
	)

	client := &http.Client{}
	req, err = http.NewRequest("POST", URL, bytes.NewReader(b))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err = client.Do(req)
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("StatusCode is %d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
