package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	RedisHost string

	DBType        string
	DBPath        string
	DBTablePrefix string
	DBName        string
	DBUser        string
	DBPassword    string

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatal("Failed to parse 'conf/app.ini:'", err)
	}

	LoadBase()
	LoadServer()
	LoadDB()
	LoadApp()
	LoadRedis()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Failed to get section 'server':", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal("Failed to get section 'app'", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadDB() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatal("Failed to get section 'database'", err)
	}
	DBType = sec.Key("TYPE").MustString("sqlite")
	DBPath = sec.Key("PATH").MustString("data/app.db")
	DBName = sec.Key("NAME").MustString("")
	DBUser = sec.Key("USER").MustString("")
	DBPassword = sec.Key("PASSWD").MustString("")
	DBTablePrefix = sec.Key("TABLE_PREFIX").MustString("wen_")
}

func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatal("Failed to get section 'redis'", err)
	}
	RedisHost = sec.Key("REDIS_HOST").MustString("")
}
