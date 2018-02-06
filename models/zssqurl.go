package models

var (
	StatisticsURL = "http://api.zhuishushenqi.com/cats/lv2/statistics"                                                     //获取所有分类
	GenderURL     = "http://api.zhuishushenqi.com/ranking/gender"                                                          //获取排行榜类型
	RankId        = "http://api.zhuishushenqi.com/ranking/%s"                                                              //获取排行榜小说
	CatsURL       = "http://api.zhuishushenqi.com/cats/lv2"                                                                //获取分类下小类型
	CategoriesURL = "https://api.zhuishushenqi.com/book/by-categories?gender=%s&type=%s&major=%s&minor=&start=%d&limit=20" //根据分类获取小说列表
	BookInfoURL   = "http://api.zhuishushenqi.com/book/%s"                                                                 //获取小说信息
	BtocURL       = "http://api.zhuishushenqi.com/btoc?view=%s&book=%s"                                                    //获取小说正版源
	AtocURL       = "http://api.zhuishushenqi.com/atoc?view=%s&book=%s"                                                    //获取小说正版源于盗版源(混合)
	MixAtocURL    = "http://api.zhuishushenqi.com/mix-atoc/%s?view=chapters"                                               //获取小说章节(根据小说id)
	SourceURL     = "http://api.zhuishushenqi.com/atoc/%s?view=chapters"                                                   //获取小说章节(根据小说源id)

	ContentURL      = "http://chapterup.zhuishushenqi.com/chapter/%s"              //获取小说章节内容
	SearchURL       = "http://api.zhuishushenqi.com/book/auto-complete?query=%s"   //搜索自动补充
	FuzzySearchURL  = "http://api.zhuishushenqi.com/book/fuzzy-search?query=%s"    //模糊搜索
	BookLastChapter = "http://api05iye5.zhuishushenqi.com/book?view=updated&id=%s" //获取小说最新章节
)
