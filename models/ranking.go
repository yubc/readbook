package models

type RankingInfo struct {
	Rankings Ranking `json:"ranking"`
	Status   bool    `json:"ok"`
}

type Ranking struct {
	ID         string  `json:"id"`
	IDs        string  `json:"_id"`
	Updated    string  `json:"updated"`
	Created    string  `json:"created"`
	Title      string  `json:"title"`
	Tag        string  `json:"tag"`
	Cover      string  `json:"cover"`
	Icon       string  `json:"icon"`
	V          int64   `json:"__v"`
	MonthRank  string  `json:"monthRank"`
	TotalRank  string  `json:"totalRank"`
	ShortTitle string  `json:"shortTitle"`
	BiTag      string  `json:"biTag"`
	IsSub      bool    `json:"isSub"`
	Collapse   bool    `json:"collapse"`
	New        bool    `json:"new"`
	Gender     string  `json:"female"`
	Priority   int64   `json:"priority"`
	Book       []Books `json:"books"`
}
type Books struct {
	ID             string `json:"_id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	ShortIntro     string `json:"shortIntro"`
	Cover          string `json:"cover"`
	Site           string `json:"site"`
	MajorCate      string `json:"majorCate"`
	MinorCate      string `json:"minorCate"`
	AllowMonthly   bool   `json:"allowMonthly"`
	Banned         int64  `json:"banned"`
	LatelyFollower int64  `json:"latelyFollower"`
	RetentionRatio string `json:"-"`
}
