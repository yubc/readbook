package models

type CategoriesList struct {
	Total  int64            `json:"total"`
	Status bool             `json:"ok"`
	Books  []CategoriesBook `json:"books"`
}
type CategoriesBook struct {
	Books
	SizeType    int64    `json:"sizetype"`
	SuperScript string   `json:"superscript"`
	ContentType string   `json:"contentType"`
	LastChapter string   `json:"lastChapter"`
	Tags        []string `json:"tags"`
}
