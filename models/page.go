package models

// a Page is like an article, but without tag
// for example you can create a about page using it
type Page struct {
	Model

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func AddPage(data map[string]interface{}) bool {
	if _, ok := data["desc"]; !ok {
		data["desc"] = ""
	}
	if _, ok := data["state"]; !ok {
		data["state"] = 0
	}
	if _, ok := data["modified_by"]; !ok {
		data["modified_by"] = data["created_by"]
	}
	db.Create(&Page{
		Title:      data["title"].(string),
		Desc:       data["desc"].(string),
		Content:    data["content"].(string),
		CreatedBy:  data["created_by"].(string),
		ModifiedBy: data["modified_by"].(string),
		State:      data["state"].(int),
	})

	return true
}

func EditPage(id int, data map[string]interface{}) bool {
	db.First(&Page{}, id).Updates(data)

	return true
}

func GetPage(id int) (page Page) {
	db.First(&page, id)

	return
}

func GetPages(maps interface{}) (pages []Page) {
	db.Where(maps).Find(&pages)
	for _, page := range pages {
		page.Content = ""
	}

	return
}

func DeletePage(id int) bool {
	db.Delete(&Page{}, id)

	return true
}

func ExistPageByID(id int) bool {
	var page Page
	db.Find(&page, id)

	return page.ID > 0
}
