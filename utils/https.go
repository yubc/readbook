package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func HttpPost(URL string, b []byte, isjson bool) (body []byte, err error) {

	var (
		resp *http.Response
		req  *http.Request
	)

	client := &http.Client{}
	req, err = http.NewRequest("POST", URL, bytes.NewReader(b))
	if err != nil {
		return
	}
	if isjson {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	}
	resp, err = client.Do(req)
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("StatusCode is %d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
