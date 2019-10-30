package models

import "fmt"

type Tag struct {
	Model

	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Icon       string `json:"icon"`
	ArticleNum int    `json:"article_num"`
	State      int    `json:"state"`
}

func GetTags(maps interface{}) (tags []Tag) {
	if err := db.Where(maps).Find(&tags).Error; err != nil {
		fmt.Println(db.Where(maps).Find(&tags))
	}
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func AddTag(name string, desc string) bool {
	db.Create(&Tag{
		Name: name,
		Desc: desc,
	})

	return true
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}
