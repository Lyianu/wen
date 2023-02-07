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
