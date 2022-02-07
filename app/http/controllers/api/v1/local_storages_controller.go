package v1

import (
	"github.com/curatorc/cngf/response"
	"github.com/gin-gonic/gin"
	"sentry-white-go/app/models/local_storage"
	"sentry-white-go/app/requests"
)

type LocalStoragesController struct {
	BaseAPIController
}

func (lsc *LocalStoragesController) Show(c *gin.Context) {
	originalModel := local_storage.Get()
	response.Data(c, originalModel)
}

func (lsc *LocalStoragesController) Update(c *gin.Context) {

	request := requests.LocalStorageRequest{}
	if ok := requests.Validate(c, &request, requests.LocalStorageSave); !ok {
		return
	}

	lsm := &local_storage.LocalStorage{
		AliyunAccessID:          request.AliyunAccessID,
		AliyunAccessSecret:      request.AliyunAccessSecret,
		AliyunAccessOSSEndpoint: request.AliyunAccessOSSEndpoint,
		AliyunAccessOSSBucket:   request.AliyunAccessOSSBucket,
	}

	isSuccess := lsm.Save()
	if !isSuccess {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	response.Data(c, lsm)
}
