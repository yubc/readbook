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

type CatsDetail struct {
	Picture []CatsList `json:"picture"`
	Male    []CatsList `json:"male"`
	Epub    []CatsList `json:"epub"`
	FeMale  []CatsList `json:"female"`
	Status  bool       `json:"ok"`
}
type CatsList struct {
	Major string   `json:"major"`
	Mins  []string `json:"mins"`
}
