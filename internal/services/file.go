package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"Go-blog-server/internal/common"
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
	"Go-blog-server/pkg/setting"
)

type FileService struct {
	dao *dao.FileDAO
}

func NewFileService() *FileService {
	return &FileService{dao: new(dao.FileDAO)}
}

// QueryFiles
func (f *FileService) UploadImg(header *multipart.FileHeader, file multipart.File) (resp common.Response, err error) {
	var (
		conf   *models.SysConf
		client *oss.Client
		bucket *oss.Bucket
	)

	content, err := ioutil.ReadAll(file)

	if err != nil {
		resp = common.Response{Err: common.ErrorReadFileFail}
		return
	}

	if conf, err = f.dao.GetOSSConf(); err != nil {
		resp = common.Response{Err: common.ErrorGetOssConfFail}
		return
	}

	if client, err = f.NewOSSClient(conf); err != nil {
		resp = common.Response{Err: common.ErrorInitOssClientFail}
		return
	}

	bucket, err = client.Bucket(setting.BucketName)

	if err != nil {
		resp = common.Response{Err: common.ErrorInitBucketFail}
	}

	key := fmt.Sprintf("%d%s", time.Now().Unix(), header.Filename)
	err = bucket.PutObject(key, bytes.NewReader(content))

	if err != nil {
		resp = common.Response{Err: common.ErrorUploadOssFail}
	}

	url := fmt.Sprintf("%s%s", setting.SourceURL, key)

	fileModel := &models.File{
		FileName: key,
		Size:     int(header.Size),
		URL:      url,
		Type:     path.Ext(key),
	}
	err = f.dao.AddFile(fileModel)

	resp = common.Response{Err: common.SUCCESS, Data: url}

	return

}

func (s *FileService) QueryFilesReq(req *models.QueryFileReq) (resp common.Response, err error) {
	var (
		total int
		files []*models.File
	)

	if total, files, err = s.dao.QueryFiles(req); err != nil {
		resp.Err = common.ErrorGetFileFali
		return
	}

	filesSerializer := models.FilesSerializer{files}

	rep := &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     filesSerializer.Response(),
	}

	resp = common.Response{Err: common.SUCCESS, Data: rep}

	return
}

func (f *FileService) NewOSSClient(conf *models.SysConf) (client *oss.Client, err error) {
	var (
		endpoint        = setting.EndPoint
		accessKeyID     = conf.OssAccessKeyID
		accessKeySecret = conf.OssAccessKeySercet
	)
	client, err = oss.New(endpoint, accessKeyID, accessKeySecret)
	return
}
