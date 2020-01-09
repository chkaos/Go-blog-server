package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"Go-blog-server/internal/common"
)

//动态连接查询
type Where func(*gorm.DB) *gorm.DB

//通用输入输出
type Out interface{}

//DAO 接口
type DaoInitializer interface {
	Insert(out Out) error
	Exec(where Where, out Out, offset int, limit int) error
	FindOne(where Where, out Out) error
	FindAll(where Where, out Out) error
	Count(where Where, out Out) (error, int)
	Delete(where Where, out Out) error
	Update(out Out, params map[string]interface{}) error
}

type BaseDaoInitializer struct {
}

//接口实现自检
var _ DaoInitializer = &BaseDaoInitializer{}

//获取数据库实例
func (bd *BaseDaoInitializer) db() *gorm.DB {
	return common.GetDB()
}

//通用查询
func (bd *BaseDaoInitializer) Exec(where Where, out Out, offset int, limit int) error {
	return bd.db().Scopes(where).Offset(offset).Limit(limit).Find(out).Error
}

//插入数据
func (bd *BaseDaoInitializer) Insert(out Out) error {
	return bd.db().Create(out).Error
}

//查单个
func (bd *BaseDaoInitializer) FindOne(where Where, out Out) error {
	fmt.Println(where)
	return bd.db().Scopes(where).First(out).Error
}

//查全部
func (bd *BaseDaoInitializer) FindAll(where Where, out Out) error {
	return bd.db().Scopes(where).Find(out).Error
}

//更新数据 ,单个 或者 多个
func (bd *BaseDaoInitializer) Update(out Out, params map[string]interface{}) error {
	return bd.db().Model(out).Update(params).Error
}

//完全更新
func (bd *BaseDaoInitializer) Save(out Out) error {
	return bd.db().Save(out).Error
}

//Count
func (bd *BaseDaoInitializer) Count(where Where, out Out) (err error, count int) {
	return bd.db().Model(out).Scopes(where).Count(&count).Error, count
}

//删除数据, 单个或者多个
func (bd *BaseDaoInitializer) Delete(where Where, out Out) error {
	return bd.db().Scopes(where).Delete(out).Error
}
