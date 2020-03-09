package common

import (
	"fmt"
	"log"
	"time"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"Go-blog-server/pkg/setting"
)

var DB *gorm.DB

func init() {
	var (
		err                                        error
		dbType, dbName, user, password, host, port string
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
	port = sec.Key("DB_PORT").String()

	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbName))

	if err != nil {
		fmt.Println("数据库连接失败")
		log.Fatal(err)
	} else {
		fmt.Println("数据库连接成功")
	}

	gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	// 全局禁用表名复数
	DB.SingularTable(true)

	DB.LogMode(setting.RunMode == "debug")

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	// 注册回调
	DB.Callback().Create().Replace("gorm:create_time_stamp", updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

}

func CloseDB() {
	DB.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedAt", time.Now().Unix())
	}
}

func GetDB() *gorm.DB {
	return DB
}

func GetDBWithTableName(tableName string) *gorm.DB {
	return DB.Table(tableName)
}
