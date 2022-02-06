package v1

import (
	"github.com/curatorc/cngf/response"
	"github.com/gin-gonic/gin"
	"sentry-white-go/app/models/original"
	"sentry-white-go/app/requests"
)

type OriginalsController struct {
	BaseAPIController
}

func (oc *OriginalsController) Index(c *gin.Context) {
	ocl := original.All()

	// 过滤 deleted
	var os []original.Original
	for _, o := range ocl.Originals {
		if o.DeletedAt.IsZero() {
			os = append(os, o)
		}
	}

	response.JSON(c, gin.H{
		"data": os,
		"meta": gin.H{
			"total": len(os),
		},
	})
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
	isSuccess := originalModel.Save()
	if !isSuccess {
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

	isSuccess := originalModel.Delete()
	if !isSuccess {
		response.Abort500(c, "删除失败，请稍后尝试~")
		return
	}

	response.Success(c)
}
