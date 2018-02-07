package models

type FuzzySearch struct {
	Total  int64       `json:"total"`
	Status bool        `json:"ok"`
	Books  []FuzzyBook `json:"books"`
}
type FuzzyBook struct {
	ID             string      `json:"_id"`
	HasCp          bool        `json:"hasCp"`
	Title          string      `json:"title"`
	Aliases        string      `json:"aliases"`
	Cat            string      `json:"cat"`
	Author         string      `json:"author"`
	Site           string      `json:"site"`
	Cover          string      `json:"cover"`
	ShortIntro     string      `json:"shortIntro"`
	LastChapter    string      `json:"lastChapter"`
	RetentionRatio interface{} `json:"retentionRatio"`
	Banned         int64       `json:"banned"`
	LatelyFollower int64       `json:"latelyFollower"`
	WordCount      int64       `json:"wordCount"`
	ContentType    string      `json:"contentType"`
	Superscript    string      `json:"superscript"`
	Sizetype       int64       `json:"sizetype"`
	Highlight      Titles      `json:"highlight"`
}

type Titles struct {
	Title []string `json:"title"`
}
