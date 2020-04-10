package services

import (
	"Go-blog-server/internal/common"
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
	"fmt"
)

type TagService struct {
	dao *dao.TagDAO
}

func NewTagService() *TagService {
	return &TagService{dao: dao.NewTagDAO()}
}

// QueryTags
func (s *TagService) QueryTagsReq(req *models.QueryTagReq) (resp common.Response, err error) {
	var (
		total int
		tags  []models.Tag
	)

	if total, tags, err = s.dao.QueryTags(req); err != nil {
		resp.Err = common.ErrorGetTagFail
		return
	}

	tagsSerializer := models.TagsSerializer{tags}

	rep := &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     tagsSerializer.Response(),
	}

	resp = common.Response{Err: common.SUCCESS, Data: rep}

	return
}

func (s *TagService) QueryAllTags() (resp common.Response, err error) {
	var tags []models.Tag

	if tags, err = s.dao.QueryAllTags(); err != nil {
		resp.Err = common.ErrorGetTagFail
		return
	}

	tagsSerializer := models.TagsSerializer{tags}
	tagRes := tagsSerializer.PreviewResponse()
	resp = common.Response{Err: common.SUCCESS, Data: tagRes}

	return
}

func (s *TagService) AddTag(tag models.Tag) (resp common.Response, err error) {

	var (
		tagModel models.Tag
	)
	tagModel, err = s.dao.QueryTagByName(tag.Name)

	if tagModel.ID > 0 {

		resp.Err = common.ErrorTagExist
		return
	}

	if err = s.dao.AddTag(tag); err != nil {
		resp.Err = common.ErrorAddTagFail
	} else {
		resp.Err = common.SUCCESS
		resp.Data = tag.PreviewResponse()
	}

	return
}

func (s *TagService) UpdateTag(tag models.Tag) (resp common.Response, err error) {

	var (
		tagModel models.Tag
	)
	tagModel, err = s.dao.QueryTagByID(tag.ID)

	fmt.Println(tagModel, err)

	if tagModel.ID == 0 {

		resp.Err = common.ErrorTagNotExist
		return
	}

	if err = s.dao.UpdateTag(tag); err != nil {
		resp.Err = common.ErrorUpdateTagFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

func (s *TagService) DeleteTag(id int) (resp common.Response, err error) {

	var (
		tagModel models.Tag
	)
	tagModel, err = s.dao.QueryTagByID(id)

	if tagModel.ID == 0 {

		resp.Err = common.ErrorTagNotExist
		return
	}

	if err = s.dao.DeleteTag(id); err != nil {
		resp.Err = common.ErrorDeleteTagFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}
