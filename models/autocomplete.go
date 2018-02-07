package models

type AutoComplete struct {
	KeyWords []string `json:"keywords"`
	Status   bool     `json:"ok"`
}
