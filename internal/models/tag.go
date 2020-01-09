package models

type TagModel struct {
	Model

	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Icon       string `json:"icon"`
	ArticleNum int    `json:"article_num"`
	State      int    `json:"state"`
}

func (TagModel) TableName() string {
	return "tag"
}

type TagResponse struct {
	ID         int `json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
	Name string  `json:"name"`
	Desc      string `json:"desc"`
	Icon   string `json:"icon"`
	ArticleNum     int    `json:"article_num"` 
	State   int    `json:"state"`
}

func (t *TagModel) TagResponse() TagResponse {
	tag := TagResponse{
		ID: t.ID,
		CreatedAt: t.CreatedAt,
		ModifiedAt: t.ModifiedAt,
		Name: t.Name,
		Desc:    t.Desc,
		Icon:    t.Icon,
		ArticleNum:    t.ArticleNum,
		State:      t.State,
	}
	return tag
}