package models

type BookInfos struct {
	ID                 string      `json:"_id"`
	Title              string      `json:"title"`
	Author             string      `json:"author"`
	LongIntro          string      `json:"longIntro"`
	Cover              string      `json:"cover"`
	Creater            string      `json:"creater"`
	MajorCate          string      `json:"majorCate"`
	MinorCate          string      `json:"minorCate"`
	HiddenPackage      []string    `json:"hiddenPackage"`
	Apptype            []int       `json:"apptype"`
	Ratings            Rating      `json:"rating"`
	HasCopyright       bool        `json:"hasCopyright"`
	Buytype            int         `json:"buytype"`
	Sizetype           int         `json:"sizetype"`
	Superscript        string      `json:"superscript"`
	Currency           int         `json:"currency"`
	ContentType        string      `json:"contentType"`
	_Le                bool        `json:"_le"`
	AllowMonthly       bool        `json:"allowMonthly"`
	AllowVoucher       bool        `json:"allowVoucher"`
	AllowBeanVoucher   bool        `json:"allowBeanVoucher"`
	HasCp              bool        `json:"hasCp"`
	PostCount          int64       `json:"postCount"`
	LatelyFollower     int64       `json:"latelyFollower"`
	FollowerCount      int64       `json:"followerCount"`
	WordCount          int64       `json:"wordCount"`
	SerializeWordCount int64       `json:"serializeWordCount"`
	RetentionRatio     string      `json:"retentionRatio"`
	Updated            string      `json:"updated"`
	IsSerial           bool        `json:"isSerial"`
	ChaptersCount      int64       `json:"chaptersCount"`
	LastChapter        string      `json:"lastChapter"`
	Gender             []string    `json:"gender"`
	Tags               []string    `json:"tags"`
	AdvertRead         bool        `json:"advertRead"`
	Cat                string      `json:"cat"`
	Donate             bool        `json:"donate"`
	_Gg                bool        `json:"_gg"`
	Discount           interface{} `json:"discount"`
	Limit              bool        `json:"limit"`
}

type Rating struct {
	Count    int     `json:"count"`
	Score    float64 `json:"score"`
	IsEffect bool    `json:"isEffect"`
}
