package v1

import (
	"github.com/curatorc/cngf/response"
	"sentry-white-go/app/models/project"
	"sentry-white-go/app/models/responsible_person"
	"sentry-white-go/app/requests"

	"github.com/gin-gonic/gin"
)

type ResponsiblePeopleController struct {
	BaseAPIController
}

func (ctrl *ResponsiblePeopleController) Index(c *gin.Context) {
	rpcl := responsible_person.All()
	// 过滤 deleted
	var rps []responsible_person.ResponsiblePerson
	for _, r := range rpcl.ResponsiblePeople {
		if r.DeletedAt.IsZero() {
			rps = append(rps, r)
		}
	}

	response.JSON(c, gin.H{
		"data": rps,
		"meta": gin.H{
			"total": len(rps),
		},
	})
}

func (ctrl *ResponsiblePeopleController) Show(c *gin.Context) {
	rm := responsible_person.Get(c.Param("id"))
	if rm.ID == 0 {
		response.Abort404(c)
		return
	}

	response.Data(c, rm)
}

func (ctrl *ResponsiblePeopleController) Store(c *gin.Context) {
	request := requests.ResponsiblePersonRequest{}
	if ok := requests.Validate(c, &request, requests.ResponsiblePersonSave); !ok {
		return
	}

	rm := responsible_person.ResponsiblePerson{
		Name:  request.Name,
		Phone: request.Phone,
	}
	rm.Create()
	if rm.ID > 0 {
		response.Created(c, rm)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ResponsiblePeopleController) Update(c *gin.Context) {

	responsiblePersonModel := responsible_person.Get(c.Param("id"))
	if responsiblePersonModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.ResponsiblePersonRequest{}
	if ok := requests.Validate(c, &request, requests.ResponsiblePersonSave); !ok {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	responsiblePersonModel.Name = request.Name
	responsiblePersonModel.Phone = request.Phone
	if ok := responsiblePersonModel.Save(); !ok {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	response.Data(c, responsiblePersonModel)
}

func (ctrl *ResponsiblePeopleController) Delete(c *gin.Context) {

	responsiblePersonModel := responsible_person.Get(c.Param("id"))
	if responsiblePersonModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if !responsiblePersonModel.Delete() {
		response.Abort500(c, "删除失败，请稍后尝试~")
		return
	}

	// 删除模型
	project.DeleteResponsiblePersonFromProject(responsiblePersonModel.ID, responsiblePersonModel.ProjectsID)
	response.Success(c)
}
