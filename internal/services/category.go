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
	return &CategoryService{dao: dao.NewCategoryDAO()}
}

// QueryCategorys
func (s *CategoryService) QueryCategorysReq(req *models.QueryCategoryReq) (resp common.Response, err error) {
	var (
		total     int
		categorys []models.Category
	)

	if total, categorys, err = s.dao.QueryCategorys(req); err != nil {
		resp.Err = common.ErrorGetCategoryFail
		return
	}

	CategorysSerializer := models.CategorysSerializer{categorys}

	rep := &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     CategorysSerializer.Response(),
	}

	resp = common.Response{Err: common.SUCCESS, Data: rep}

	return
}

// QueryAllCategorys
func (s *CategoryService) QueryAllCategorys() (resp common.Response, err error) {
	var Categorys []models.Category

	if Categorys, err = s.dao.QueryAllCategorys(); err != nil {
		resp.Err = common.ErrorGetCategoryFail
		return
	}

	CategorysSerializer := models.CategorysSerializer{Categorys}
	CategoryRes := CategorysSerializer.PreviewResponse()
	resp = common.Response{Err: common.SUCCESS, Data: CategoryRes}

	return
}

// AddCategory
func (s *CategoryService) AddCategory(Category models.Category) (resp common.Response, err error) {

	var (
		CategoryModel models.Category
	)
	CategoryModel, err = s.dao.QueryCategoryByName(Category.Name)

	if CategoryModel.ID > 0 {

		resp.Err = common.ErrorCategoryExist
		return
	}

	if err = s.dao.AddCategory(Category); err != nil {
		resp.Err = common.ErrorAddCategoryFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

// UpdateCategory
func (s *CategoryService) UpdateCategory(Category models.Category) (resp common.Response, err error) {

	var (
		CategoryModel models.Category
	)
	CategoryModel, err = s.dao.QueryCategoryByID(Category.ID)

	fmt.Println(CategoryModel, err)

	if CategoryModel.ID == 0 {

		resp.Err = common.ErrorCategoryNotExist
		return
	}

	if err = s.dao.UpdateCategory(Category); err != nil {
		resp.Err = common.ErrorUpdateCategoryFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

// DeleteCategory
func (s *CategoryService) DeleteCategory(id int) (resp common.Response, err error) {

	var (
		CategoryModel models.Category
	)
	CategoryModel, err = s.dao.QueryCategoryByID(id)

	if CategoryModel.ID == 0 {
		resp.Err = common.ErrorCategoryNotExist
		return
	}

	if err = s.dao.DeleteCategory(id); err != nil {
		resp.Err = common.ErrorDeleteCategoryFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}
