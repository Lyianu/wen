package models

import (
	"log"

	"github.com/Lyianu/wen/pkg/setting"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type Model struct {
	gorm.Model

	ID         uint `gorm:"primaryKey" json:"id"`
	CreatedAt  int  `json:"created_at"`
	ModifiedAt int  `json:"modified_at"`
}

func init() {
	var err error
	var dbPath string

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal("Failed to load section 'database'", err)
	}

	dbPath = sec.Key("PATH").String()
	// enable temporary memory database
	if setting.RunMode == "debug" {
		dbPath = "file::memory:?cache=shared"
	}
	tablePrefix := sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&Tag{}, &Article{}, &Auth{}, &Page{}, &Site{})

	if setting.RunMode == "debug" {
		// No copyright picture
		AddSite("Wen", "https://www.pexels.com/zh-cn/photo/13258051/https://images.pexels.com/photos/13258051/pexels-photo-13258051.jpeg?cs=srgb&dl=pexels-alteredsnaps-13258051.jpg&fm=jpg&w=1920&h=1282")
		AddAuth("test", "123")
		AddPage(map[string]interface{}{"title": "About", "content": "Lorem", "desc": "lorem", "created_by": "Wen-authors"})
	}
	//AddAuth("test", "test123")
}
