package models

import (
	"fmt"
	"log"

	"github.com/Lyianu/wen/pkg/setting"
	"gorm.io/driver/mysql"
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

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal("Failed to load section 'database'", err)
	}

	// enable temporary memory database
	if setting.RunMode == "debug" {
		setting.DBPath = "file::memory:?cache=shared"
		setting.DBType = "sqlite3"
	}
	tablePrefix := sec.Key("TABLE_PREFIX").String()

	if setting.DBType == "sqlite3" {
		db, err = gorm.Open(sqlite.Open(setting.DBPath), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   tablePrefix,
				SingularTable: true,
			},
		})
		if err != nil {
			panic("failed to connect to database")
		}
	}

	if setting.DBType == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", setting.DBUser, setting.DBPassword, setting.DBPath, setting.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   tablePrefix,
				SingularTable: true,
			},
		})
		if err != nil {
			panic("failed to connect to database")
		}
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
