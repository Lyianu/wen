package models

type Article struct {
	Model

	TagID []int  `json:"tag_id" gorm:"-"`
	Tags  []*Tag `json:"tags" gorm:"many2many:article_tags"`

	Page
}

func ExistArticleByID(id int) bool {
	var article Article
	db.First(&article, id)
	if int(article.ID) == id {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) int {
	var count int64
	db.Model(&Article{}).Where(maps).Count(&count)
	return int(count)
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tags").Where(maps).Order("id DESC").Offset(pageNum).Limit(pageSize).Find(&articles)
	for _, v := range articles {
		v.Content = ""
	}
	return
}

func GetArticle(id int) (article Article) {
	db.First(&article, id)

	return
}

func EditArticle(id int, data interface{}) bool {
	db.First(&Article{}, id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	tags := FindTags(data["tag_id"].([]int)...)
	db.Create(&Article{
		TagID: data["tag_id"].([]int),
		Tags:  tags,
		Page: Page{
			Title:     data["title"].(string),
			Desc:      data["desc"].(string),
			Content:   data["content"].(string),
			CreatedBy: data["created_by"].(string),
			State:     data["state"].(int),
		},
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Delete(&Article{}, id)

	return true
}
