package models

type Statistics struct {
	Male   []BookInfo `json:"male"`
	Status bool       `json:"ok"`
}

type BookInfo struct {
	Name         string   `json:"name"`
	BookCount    int      `json:"bookCount"`
	MonthlyCount int      `json:"monthlyCount"`
	Icon         string   `json:"icon"`
	BookCover    []string `json:"bookCover"`
}
