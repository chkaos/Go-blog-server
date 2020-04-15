package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg    *ini.File
	EnvCfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	TimeFormat string
	JwtSecret  string

	EndPoint   string
	BucketName string
	SourceURL  string
)

func init() {
	var (
		envCfgPath string
		err        error
	)
	Cfg, err = ini.Load("./conf/base.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/base.ini': %v", err)
	}

	LoadBase()

	if RunMode == "debug" {
		envCfgPath = "dev.ini"
	} else {
		envCfgPath = "prod.ini"
	}

	EnvCfg, err = ini.Load(fmt.Sprintf("./conf/%s", envCfgPath))

	if err != nil {
		log.Fatalf("Fail to parse 'conf/%s': %v", envCfgPath, err)
	}

	LoadOSS()
	LoadServer()
	LoadApp()

}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	TimeFormat = sec.Key("TIME_FORMAT").MustString("")
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}

func LoadOSS() {
	sec, err := Cfg.GetSection("oss")
	if err != nil {
		log.Fatalf("Fail to get section 'oss': %v", err)
	}

	EndPoint = sec.Key("END_POINT").MustString("")
	BucketName = sec.Key("BUCKET_NAME").MustString("")
	SourceURL = sec.Key("SOURSE_URL").MustString("")

}
