package services

import (
	"Go-blog-server/internal/common"
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
	"fmt"
)

type CategoryService struct {
	dao *dao.CategoryDAO
}

func NewCategoryService() *CategoryService {
	return &CategoryService{dao: new(dao.CategoryDAO)}
}

// QueryCategorys
func (s *CategoryService) QueryCategorysReq(req *models.QueryCategoryReq) (resp common.Response, err error) {
	var (
		total     int
		Categorys []*models.Category
	)

	if total, Categorys, err = s.dao.QueryCategorys(req); err != nil {
		resp.Err = common.ERROR_GET_CATEGORY_FAIL
		return
	}

	CategorysSerializer := models.CategorysSerializer{Categorys}

	rep := &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     CategorysSerializer.Response(),
	}

	resp = common.Response{Err: common.SUCCESS, Data: rep}

	return
}

func (s *CategoryService) QueryAllCategorys() (resp common.Response, err error) {
	var Categorys []*models.Category

	if Categorys, err = s.dao.QueryAllCategorys(); err != nil {
		resp.Err = common.ERROR_GET_CATEGORY_FAIL
		return
	}

	CategorysSerializer := models.CategorysSerializer{Categorys}
	CategoryRes := CategorysSerializer.PreviewResponse()
	resp = common.Response{Err: common.SUCCESS, Data: CategoryRes}

	return
}

func (s *CategoryService) AddCategory(Category *models.Category) (resp common.Response, err error) {

	var (
		CategoryModel *models.Category
	)
	CategoryModel, err = s.dao.QueryCategoryByName(Category.Name)

	if CategoryModel.ID > 0 {

		resp.Err = common.ERROR_CATEGORY_EXIST
		return
	}

	if err = s.dao.AddCategory(Category); err != nil {
		resp.Err = common.ERROR_ADD_CATEGORY_FAIL
	} else {
		resp.Err = common.SUCCESS
		resp.Data = Category.PreviewResponse()
	}

	return
}

func (s *CategoryService) UpdateCategory(Category *models.Category) (resp common.Response, err error) {

	var (
		CategoryModel *models.Category
	)
	CategoryModel, err = s.dao.QueryCategoryByID(Category.ID)

	fmt.Println(CategoryModel, err)

	if CategoryModel.ID == 0 {

		resp.Err = common.ERROR_CATEGORY_NOT_EXIST
		return
	}

	if err = s.dao.UptadeCategory(Category); err != nil {
		resp.Err = common.ERROR_UPDATE_CATEGORY_FAIL
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

func (s *CategoryService) DeleteCategory(id int) (resp common.Response, err error) {

	var (
		CategoryModel *models.Category
	)
	CategoryModel, err = s.dao.QueryCategoryByID(id)

	if CategoryModel.ID == 0 {

		resp.Err = common.ERROR_CATEGORY_NOT_EXIST
		return
	}

	if err = s.dao.DeleteCategory(id); err != nil {
		resp.Err = common.ERROR_DETELE_CATEGORY_FAIL
	} else {
		resp.Err = common.SUCCESS
	}

	return
}
