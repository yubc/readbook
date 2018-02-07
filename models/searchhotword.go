package models

type SearchHotwordError struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Status bool   `json:"ok"`
}
