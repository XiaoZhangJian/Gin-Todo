package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	RunMode      string
	JwtSecret    string
	PageSize     int
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("fail to read file %v \n", err)
	}
	LoadBase()
	loadServe()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServe() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("fail to get section 'server' : %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	ReadTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app' : %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@#$!@#$!@#$!@#$")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
