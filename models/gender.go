package models

type Gender struct {
	Picture []GenderList `json:"picture"`
	Male    []GenderList `json:"male"`
	Epub    []GenderList `json:"epub"`
	FeMale  []GenderList `json:"female"`
	Status  bool         `json:"ok"`
}

type GenderList struct {
	Id         string `json:"_id"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	Collapse   bool   `json:"collapse"`
	MonthRank  string `json:"monthRank"`
	TotalRank  string `json:"totalRank"`
	ShortTitle string `json:"shortTitle"`
}
