package validators

import "Go-blog-server/internal/models"

type AddBulletinForm struct {
	Content string `form:"content" valid:"Required;MaxSize(100)"`
	Source  string `form:"source" valid:"MaxSize(20)"`
	Top     int    `form:"top" valid:"Range(0,1)"`
}

type EditBulletinForm struct {
	ID      int    `form:"id" valid:"Required;Min(1)"`
	Content string `form:"content" valid:"Required;MaxSize(100)"`
	Source  string `form:"source" valid:"MaxSize(20)"`
	Top     int    `form:"top" valid:"Range(0,1)"`
}

func (a *AddBulletinForm) Transform() (bulletin models.Bulletin) {
	return models.Bulletin{
		Content: a.Content,
		Source:  a.Source,
		Top:     a.Top,
	}
}

func (e *EditBulletinForm) Transform() (bulletin models.Bulletin) {
	bulletin = models.Bulletin{
		Content: e.Content,
		Source:  e.Source,
		Top:     e.Top,
	}
	bulletin.ID = e.ID
	return
}
