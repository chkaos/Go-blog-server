package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"Go-blog-server/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("DB_TYPE").String()
	dbName = sec.Key("DB_NAME").String()
	user = sec.Key("DB_USER").String()
	password = sec.Key("DB_PASSWORD").String()
	host = sec.Key("DB_HOST").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	// 全局禁用表名复数
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 注册回调
	db.Callback().Create().Register("my_plugin:before_create", beforeCreate)
	db.Callback().Update().Before("gorm:update").Register("my_plugin:before_update", beforeUpdate)
}

func CloseDB() {
	defer db.Close()
}

func beforeCreate(scope *gorm.Scope) {
	if scope.HasColumn("created_on") {
		scope.SetColumn("created_on", time.Now().Unix())
	}
}

func beforeUpdate(scope *gorm.Scope) {
	if scope.HasColumn("modified_on") {
		scope.SetColumn("modified_on", time.Now().Unix())
	}
}
