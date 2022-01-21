package v1

import (
	"github.com/curatorc/cngf/response"
	"sentry-white-go/app/models/original"
	"sentry-white-go/app/requests"

	"github.com/gin-gonic/gin"
)

type OriginalsController struct {
	BaseAPIController
}

func (oc *OriginalsController) Index(c *gin.Context) {
	originals := original.All()
	response.Data(c, originals)
}

func (oc *OriginalsController) Show(c *gin.Context) {
	originalModel := original.Get(c.Param("id"))
	if originalModel.ID == 0 {
		response.Abort404(c)
		return
	}

	response.Data(c, originalModel)
}

func (oc *OriginalsController) Store(c *gin.Context) {

	request := requests.OriginalRequest{}
	if ok := requests.Validate(c, &request, requests.OriginalSave); !ok {
		return
	}

	originalModel := original.Original{
		Name:        request.Name,
		AccountName: request.AccountName,
		Password:    request.Password,
	}
	originalModel.Create()
	if originalModel.ID > 0 {
		response.Created(c, originalModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (oc *OriginalsController) Update(c *gin.Context) {

	originalModel := original.Get(c.Param("id"))
	if originalModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.OriginalRequest{}
	if ok := requests.Validate(c, &request, requests.OriginalSave); !ok {
		return
	}

	originalModel.Name = request.Name
	originalModel.AccountName = request.AccountName
	originalModel.Password = request.Password
	rowsAffected := originalModel.Save()
	if rowsAffected != 1 {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	response.Data(c, originalModel)
}

func (oc *OriginalsController) Delete(c *gin.Context) {

	originalModel := original.Get(c.Param("id"))
	if originalModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := originalModel.Delete()
	if rowsAffected != 1 {
		response.Abort500(c, "删除失败，请稍后尝试~")
		return
	}

	response.Success(c)
}
