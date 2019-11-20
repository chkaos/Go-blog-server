package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model

	CategoryID      int    `json:"category_id"`
	Tags            []Tag  `gorm:"many2many:tag;"`
	Title           string `json:"title"`
	Desc            string `json:"desc"`
	Keywords        string `json:"keywords"`
	Content         string `json:"content"`
	RenderedContent string `json:"rendered_content"`
	Thumb           string `json:"thumb"`
	Source          int    `json:"source"`
	ReproduceURL    string `json:"reproduce_url"`
	State           int    `json:"state"`
	Likenum         int    `json:"like_num"`
	Commentsnum     int    `json:"comments_num"`
	Pvsnum          int    `json:"pvs_num"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }

	return &article, nil
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		CategoryID:      data["category_id"].(int),
		Title:           data["title"].(string),
		Desc:            data["desc"].(string),
		Content:         data["content"].(string),
		Keywords:        data["keywords"].(string),
		RenderedContent: data["RenderedContent"].(string),
		Thumb:           data["thumb"].(string),
		Source:          data["source"].(int),
		ReproduceURL:    data["reproduce_url"].(string),
		State:           data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}
