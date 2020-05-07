package services

import (
	"Go-blog-server/internal/common"
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
)

type BulletinService struct {
	dao *dao.BulletinDAO
}

func NewBulletinService() *BulletinService {
	return &BulletinService{dao: dao.NewBulletinDAO()}
}

// QueryBulletins
func (s *BulletinService) QueryBulletinsReq(req *models.QueryBulletinReq) (resp common.Response, err error) {
	var (
		total     int
		bulletins []models.Bulletin
	)

	if total, bulletins, err = s.dao.QueryBulletins(req); err != nil {
		resp.Err = common.ErrorGetBulletinFail
		return
	}

	BulletinsSerializer := models.BulletinsSerializer{bulletins}

	rep := &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     BulletinsSerializer.Response(),
	}

	resp = common.Response{Err: common.SUCCESS, Data: rep}

	return
}

// AddBulletin
func (s *BulletinService) AddBulletin(Bulletin models.Bulletin) (resp common.Response, err error) {

	if err = s.dao.AddBulletin(Bulletin); err != nil {
		resp.Err = common.ErrorAddBulletinFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

// UpdateBulletin
func (s *BulletinService) UpdateBulletin(Bulletin models.Bulletin) (resp common.Response, err error) {

	var (
		BulletinModel models.Bulletin
	)
	BulletinModel, err = s.dao.QueryBulletinByID(Bulletin.ID)

	if BulletinModel.ID == 0 {

		resp.Err = common.ErrorBulletinNotExist
		return
	}

	if err = s.dao.UpdateBulletin(Bulletin); err != nil {
		resp.Err = common.ErrorUpdateBulletinFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

// DeleteBulletin
func (s *BulletinService) DeleteBulletin(id int) (resp common.Response, err error) {

	var (
		BulletinModel models.Bulletin
	)
	BulletinModel, err = s.dao.QueryBulletinByID(id)

	if BulletinModel.ID == 0 {
		resp.Err = common.ErrorBulletinNotExist
		return
	}

	if err = s.dao.DeleteBulletin(id); err != nil {
		resp.Err = common.ErrorDeleteBulletinFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}
