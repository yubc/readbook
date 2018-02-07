package models

type ChapterInfo struct {
	ChapterList []Chapters `json:"chapters"`
	Updated     string     `json:"updated"`
	Host        string     `json:"host"`
	ID          string     `json:"_id"`
	Source      string     `json:"source"`
	Name        string     `json:"name"`
	Link        string     `json:"link"`
	Book        string     `json:"book"`
}

type Chapters struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	ID        string `json:"id"`
	TotalPage int64  `json:"totalpage"`
	Partsize  int64  `json:"partsize"`
	Order     int64  `json:"order"`
	Currency  int64  `json:"currency"`
	Unreadble bool   `json:"unreadble"`
	IsVip     bool   `json:"isVip"`
}

type ChapterContentInfo struct {
	Chapter ChapterContent `json:"chapter"`
	Status  bool           `json:"ok"`
}
type ChapterContent struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	IsVip     bool   `json:"isVip"`
	CpContent string `json:"cpContent"`
	Currency  int64  `json:"currency"`
	ID        string `json:"id"`
}
