package v1

import (
	"github.com/curatorc/cngf/helpers"
	"github.com/curatorc/cngf/response"
	"sentry-white-go/app/models/pivots"
	"sentry-white-go/app/models/project"
	"sentry-white-go/app/requests"

	"github.com/gin-gonic/gin"
)

type ProjectsController struct {
	BaseAPIController
}

func (ctrl *ProjectsController) Index(c *gin.Context) {
	pcl := project.All()
	// 过滤 deleted
	var ps []project.Project
	for _, p := range pcl.Projects {
		if p.DeletedAt.IsZero() {
			ps = append(ps, p)
		}
	}

	response.JSON(c, gin.H{
		"data": ps,
		"meta": gin.H{
			"total": len(ps),
		},
	})
}

func (ctrl *ProjectsController) Show(c *gin.Context) {
	pm := project.Get(c.Param("id"))
	if pm.ID == 0 {
		response.Abort404(c)
		return
	}

	response.Data(c, pm)
}

func (ctrl *ProjectsController) Store(c *gin.Context) {

	request := requests.ProjectRequest{}
	if ok := requests.Validate(c, &request, requests.ProjectSave); !ok {
		return
	}

	pm := project.Project{
		Name:                request.Name,
		SubstituteName:      request.SubstituteName,
		OriginalID:          request.OriginalID,
		ResponsiblePeopleID: request.ResponsiblePeopleID,
	}

	pm.Create()
	if pm.ID > 0 {

		// 绑定负责人
		for _, rpID := range pm.ResponsiblePeopleID {
			pivots.BindProjectResponsiblePerson(rpID, pm.ID, pm.ResponsiblePeopleID)
		}

		response.Created(c, pm)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ProjectsController) Update(c *gin.Context) {

	projectModel := project.Get(c.Param("id"))
	oldOriginalID := projectModel.OriginalID
	oldResponsiblePeopleID := projectModel.ResponsiblePeopleID
	if projectModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.ProjectRequest{}
	if ok := requests.Validate(c, &request, requests.ProjectSave); !ok {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	projectModel.Name = request.Name
	projectModel.SubstituteName = request.SubstituteName
	projectModel.OriginalID = request.OriginalID
	projectModel.ResponsiblePeopleID = request.ResponsiblePeopleID
	if ok := projectModel.Save(oldOriginalID); !ok {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	// 更新负责人
	// 遍历新负责人，不存在的加
	for _, rpID := range request.ResponsiblePeopleID {
		if !helpers.Uint64InSlice(rpID, oldResponsiblePeopleID) {
			pivots.BindProjectResponsiblePerson(rpID, projectModel.ID, projectModel.ResponsiblePeopleID)
		}
	}
	// 遍历老负责人，不存在的减
	for _, rpID := range oldResponsiblePeopleID {
		if !helpers.Uint64InSlice(rpID, request.ResponsiblePeopleID) {
			pivots.UnBindProjectResponsiblePerson(rpID, projectModel.ID, projectModel.ResponsiblePeopleID)
		}
	}

	response.Data(c, projectModel)
}

func (ctrl *ProjectsController) Delete(c *gin.Context) {

	projectModel := project.Get(c.Param("id"))
	if projectModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := projectModel.Delete(); !ok {
		response.Abort500(c, "删除失败，请稍后尝试~")
		return
	}

	response.Success(c)
}
