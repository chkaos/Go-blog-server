package validators

import "Go-blog-server/internal/models"

type AddTagForm struct {
	Name string `form:"name" valid:"Required;MaxSize(100)"`
	Desc string `form:"desc" valid:"Required;MaxSize(100)"`
	Icon string `form:"string" valid:"MaxSize(100)"`
}

type EditTagForm struct {
	ID   int    `form:"id" valid:"Required;Min(1)"`
	Name string `form:"name" valid:"Required;MaxSize(100)"`
	Desc string `form:"desc" valid:"Required;MaxSize(100)"`
	Icon string `form:"string" valid:"MaxSize(100)"`
}

func (a *AddTagForm) Transform() (tag models.Tag) {
	return models.Tag{
		Name: a.Name,
		Desc: a.Desc,
		Icon: a.Icon,
	}
}

func (e *EditTagForm) Transform() (tag models.Tag) {
	tag = models.Tag{
		Name: e.Name,
		Desc: e.Desc,
		Icon: e.Icon,
	}
	tag.ID = e.ID
	return
}
