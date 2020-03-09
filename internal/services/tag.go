package services

import (
	"fmt"

	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
	"Go-blog-server/pkg/e"
)

type TagService struct {
	dao *dao.TagDAO
}

func NewTagService() *TagService {
	return &TagService{dao: new(dao.TagDAO)}
}

// QueryTags 
func (s *TagService) QueryTagsReq(req *models.QueryTagReq) (rep *models.PaginationRep, err error) {
	var (
		total                int
		tags []*models.Tag
	)

	if total, tags, err = s.dao.QueryTags(req); err != nil {
		return
	}

	tagsSerializer := models.TagsSerializer{tags}

	rep = &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List: tagsSerializer.Response(),
	}
	return
}

func (s *TagService) QueryAllTags() (res []models.TagResponse, err error) {
	var tags []*models.Tag

	if tags, err = s.dao.QueryAllTags(); err != nil {
		return
	}

	tagsSerializer := models.TagsSerializer{tags}
	res = tagsSerializer.PreviewResponse()

	return 
}

func (s *TagService) AddTag(tag *models.Tag) (error) {

  var (
		res *models.Tag
		err error
	)
	res, err = s.dao.QueryTagByName(tag.Name)

	if res.ID == 0 {
		err = s.dao.AddTag(tag)
	} else {
    err = e.DataExisted
	}
		
	return  err
}
