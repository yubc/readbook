package models

type SummaryList struct {
	ID            string `json:"_id"`
	LastChapter   string `json:"lastChapter`
	Link          string `json:"link"`
	Source        string `json:"source"`
	Name          string `json:"name"`
	IsCharge      bool   `json:"isCharge"`
	ChaptersCount int64  `json:"chaptersCount"`
	Updated       string `json:"updated"`
	Starting      bool   `json:"starting"`
	Host          string `json:"host"`
}
