package common

import (
	"fmt"
)

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err Err) Error() string {
	return fmt.Sprintf("ErrCode:%d , ErrMsg:%s", err.Code, err.Msg)
}

//返回结果中的错误码表示了用户调用API 的结果其中，code 为公共错误码，其适用于所有模块的 API 接口
//若 code 为 0，表示调用成功，否则，表示调用失败当调用失败后，用户可以根据下表确定错误原因并采取相应措施
var (
	SUCCESS                = Err{Code: 200, Msg: "成功"}
	ERROR                  = Err{Code: 500, Msg: "内部错误"}
	ErrorInvalidParams     = Err{Code: 4000, Msg: "缺少必要参数，或者参数值/路径格式不正确"}
	ErrorGetOssConfFail    = Err{Code: 5001, Msg: "OSS配置不存在或获取失败"}
	ErrorInitOssClientFail = Err{Code: 5002, Msg: "OSS对象生成失败"}
	ErrorUploadOssFail     = Err{Code: 5003, Msg: "OSS上传文件失败"}
	ErrorInitBucketFail    = Err{Code: 5003, Msg: "OSS获取BUCKET失败"}
	ErrorReadFileFail      = Err{Code: 5004, Msg: "读取文件失败"}
	ErrAuthorized          = Err{Code: 4100, Msg: "签名鉴权失败，请参考文档中鉴权部分"}
	ErrAccoutDeny          = Err{Code: 4200, Msg: "帐号被封禁，或者不在接口针对的用户范围内等"}
	ErrNotFound            = Err{Code: 4300, Msg: "资源不存在，或者访问了其他用户的资源"}
	ErrMethodNotAllow      = Err{Code: 4400, Msg: "协议不支持，请参考文档说明"}
	ErrSignParams          = Err{Code: 4500, Msg: "签名错误，请参考文档说明"}
	ErrInternal            = Err{Code: 5000, Msg: "服务器内部出现错误，请稍后重试"}
)

//模块错误码
var (
	ErrorAddTagSuccess = Err{Code: 0, Msg: "标签添加成功"}
	ErrorAddTagFail    = Err{Code: 10200, Msg: "标签添加失败"}
	ErrorTagExist      = Err{Code: 10201, Msg: "标签已存在"}
	ErrorGetTagFail    = Err{Code: 10202, Msg: "获取标签失败"}
	ErrorDeleteTagFail = Err{Code: 10203, Msg: "删除标签失败"}
	ErrorTagNotExist   = Err{Code: 10204, Msg: "标签不存在"}
	ErrorUpdateTagFail = Err{Code: 10205, Msg: "标签更新失败"}

	ErrorAddCategorySuccess = Err{Code: 0, Msg: "分类添加成功"}
	ErrorAddCategoryFail    = Err{Code: 10210, Msg: "分类添加失败"}
	ErrorCategoryExist      = Err{Code: 10211, Msg: "分类已存在"}
	ErrorGetCategoryFail    = Err{Code: 10212, Msg: "获取分类失败"}
	ErrorDeleteCategoryFail = Err{Code: 10213, Msg: "删除分类失败"}
	ErrorCategoryNotExist   = Err{Code: 10214, Msg: "分类不存在"}
	ErrorUpdateCategoryFail = Err{Code: 10215, Msg: "分类更新失败"}

	ErrorAddFileSuccess = Err{Code: 0, Msg: "文件上传成功"}
	ErrorAddFileFail    = Err{Code: 10220, Msg: "文件上传失败"}
	ErrorGetFileFail    = Err{Code: 10212, Msg: "获取文件失败"}

	ErrorAddArticleSuccess = Err{Code: 0, Msg: "文章添加成功"}
	ErrorAddArticleFail    = Err{Code: 10230, Msg: "文章添加失败"}
	ErrorArticleExist      = Err{Code: 10231, Msg: "文章已存在"}
	ErrorGetArticleFail    = Err{Code: 10232, Msg: "获取文章失败"}
	ErrorDeleteArticleFail = Err{Code: 10233, Msg: "删除文章失败"}
	ErrorArticleNotExist   = Err{Code: 10234, Msg: "文章不存在"}
	ErrorUpdateArticleFail = Err{Code: 10235, Msg: "文章更新失败"}

	ErrAuth        = Err{Code: 10300, Msg: "用户名或密码错误"}
	ErrorAuthToken = Err{Code: 10301, Msg: "tokn生成失败"}
)
