package validators

import "Go-blog-server/internal/models"

type AddCategoryForm struct {
	Name string `form:"name" valid:"Required;MaxSize(100)"`
	Desc string `form:"desc" valid:"Required;MaxSize(100)"`
	Icon string `form:"string" valid:"MaxSize(100)"`
}

type EditCategoryForm struct {
	ID   int    `form:"id" valid:"Required;Min(1)"`
	Name string `form:"name" valid:"Required;MaxSize(100)"`
	Desc string `form:"desc" valid:"Required;MaxSize(100)"`
	Icon string `form:"string" valid:"MaxSize(100)"`
}

func (a *AddCategoryForm) Transform() (tag models.Category) {
	return models.Category{
		Name: a.Name,
		Desc: a.Desc,
	}
}

func (e *EditCategoryForm) Transform() (tag models.Category) {
	tag = models.Category{
		Name: e.Name,
		Desc: e.Desc,
	}
	tag.ID = e.ID
	return
}
