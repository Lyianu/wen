package models

type Tag struct {
	Model

	Articles []*Article `gorm:"many2many:article_tags"`

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) int {
	var count64 int64
	db.Model(&Tag{}).Where(maps).Count(&count64)

	return int(count64)
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.First(&tag, id)
	if int(tag.ID) == id {
		return true
	}

	return false
}

// ExistTagsByID checks every Tag specified by ids
// it returns true only when every Tag exists
func ExistTagsByID(ids ...int) bool {
	for _, id := range ids {
		if !ExistTagByID(id) {
			return false
		}
	}
	return true
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

func FindTags(ids ...int) (tags []*Tag) {
	for _, id := range ids {
		var tag *Tag
		db.First(tag, id)
		tags = append(tags, tag)
	}
	return
}
