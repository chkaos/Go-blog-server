package models

import "fmt"

type Category struct {
	Model

	Name string `json:"name"`
	Desc string `json:"desc"`
}

func GetCategorys(maps interface{}) (categorys []Category) {
	if err := db.Where(maps).Find(&categorys).Error; err != nil {
		fmt.Println(db.Where(maps).Find(&categorys))
	}
	return
}

func GetCategorysTotal(maps interface{}) (count int) {
	db.Model(&Category{}).Where(maps).Count(&count)

	return
}

func AddCategory(name string, desc string) bool {
	db.Create(&Category{
		Name: name,
		Desc: desc,
	})

	return true
}

func ExistCategoryByName(name string) bool {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return true
	}

	return false
}

func ExistCategoryByID(id int) bool {
	var category Category
	db.Select("id").Where("id = ?", id).First(&category)
	if category.ID > 0 {
		return true
	}

	return false
}

func DeleteCategory(id int) bool {
	db.Where("id = ?", id).Delete(&Category{})

	return true
}

func EditCategory(id int, data interface{}) bool {
	db.Model(&Category{}).Where("id = ?", id).Updates(data)

	return true
}
