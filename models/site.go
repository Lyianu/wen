package models

// Site represents infomation of the website
// For example, Name and Hostname, etc.
type Site struct {
	Model

	SiteImageURL string `json:"image_url"` // hero image
	Name         string `json:"name"`      // site name
	BgTitle      string `json:"bg_title"`  // background title
	Desc         string `json:"desc"`      //site description
}

// Variable Wen Contains the current config of the website
var Wen Site

// As a website only has one config, AddSite should be called only once
func AddSite(name, img_url, bgTitle, desc string) bool {
	db.Create(&Site{
		Name:         name,
		SiteImageURL: img_url,
		BgTitle:      bgTitle,
		Desc:         desc,
	})

	UpdateSite()
	return true
}

// EditSite should not be call before AddSite as it does not handle error
func EditSite(data map[string]interface{}) bool {
	db.First(&Site{}, 1).Updates(data)

	UpdateSite()
	return true
}

// UpdateSite updates the site config stored in main memory
func UpdateSite() bool {
	db.First(&Wen)

	return true
}
