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

	db.AutoMigrate(&Tag{}, &Article{}, &User{}, &Page{}, &Site{})

	if setting.RunMode == "debug" {
		// No copyright picture
		//AddSite("Wen", "/pexels-flo-dahm-529643.jpg", "Lorem", "Wen blogging platform")
		//AddAuth("test", "123")
		AddPage(map[string]interface{}{"title": "About", "content": "<p>This site is built with gin as backend and react as frontend</p>", "desc": "lorem", "created_by": "Wen-authors"})
		AddPage(map[string]interface{}{"title": "Contact", "content": "Lorem l<h1>ipsum</h1>", "desc": "lorem", "created_by": "Wen-authors"})
	}
	//AddAuth("test", "test123")
}
