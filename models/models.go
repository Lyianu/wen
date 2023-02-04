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
	ID         uint `gorm:"primaryKey" json:"id"`
	CreatedOn  int  `json:"created_on"`
	ModifiedOn int  `json:"modified_on"`
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

}
