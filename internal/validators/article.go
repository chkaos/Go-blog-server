package validators

type AddArticleForm struct {
	Name  string `form:"name" valid:"Required;MaxSize(100)"`
	Desc  string `form:"desc" valid:"Required;MaxSize(100)"`
	Icon  string `form:"string" valid:"MaxSize(100)"`
	State int    `form:"state" valid:"Range(0,1)"`
}

type EditArticleForm struct {
	Id    uint   `form:"name" valid:"Required;Min(1)"`
	Name  string `form:"name" valid:"Required;MaxSize(100)"`
	Desc  string `form:"desc" valid:"Required;MaxSize(100)"`
	Icon  string `form:"string" valid:"MaxSize(100)"`
	State int    `form:"state" valid:"Range(0,1)"`
}
