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
	DeletedAt  int `json:"deleted_at"`
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
	db.Callback().Create().Replace("gorm:create_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// func deleteCallback(scope *gorm.Scope) {
//     if !scope.HasError() {
//         var extraOption string
//         if str, ok := scope.Get("gorm:delete_option"); ok {
//             extraOption = fmt.Sprint(str)
//         }

//         deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

//         if !scope.Search.Unscoped && hasDeletedOnField {
//             scope.Raw(fmt.Sprintf(
//                 "UPDATE %v SET %v=%v%v%v",
//                 scope.QuotedTableName(),
//                 scope.Quote(deletedOnField.DBName),
//                 scope.AddToVars(time.Now().Unix()),
//                 addExtraSpaceIfExist(scope.CombinedConditionSql()),
//                 addExtraSpaceIfExist(extraOption),
//             )).Exec()
//         } else {
//             scope.Raw(fmt.Sprintf(
//                 "DELETE FROM %v%v%v",
//                 scope.QuotedTableName(),
//                 addExtraSpaceIfExist(scope.CombinedConditionSql()),
//                 addExtraSpaceIfExist(extraOption),
//             )).Exec()
//         }
//     }
// }

// func addExtraSpaceIfExist(str string) string {
//     if str != "" {
//         return " " + str
//     }
//     return ""
// }
