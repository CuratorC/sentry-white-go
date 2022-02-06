package v1

import (
	"github.com/curatorc/cngf/response"
	"sentry-white-go/app/models/robot"
	"sentry-white-go/app/requests"

	"github.com/gin-gonic/gin"
)

type RobotsController struct {
	BaseAPIController
}

func (ctrl *RobotsController) Index(c *gin.Context) {
	rcl := robot.All()
	// 过滤 deleted
	var rs []robot.Robot
	for _, r := range rcl.Robots {
		if r.DeletedAt.IsZero() {
			rs = append(rs, r)
		}
	}

	response.JSON(c, gin.H{
		"data": rs,
		"meta": gin.H{
			"total": len(rs),
		},
	})
}

func (ctrl *RobotsController) Show(c *gin.Context) {
	rm := robot.Get(c.Param("id"))
	if rm.ID == 0 {
		response.Abort404(c)
		return
	}

	response.Data(c, rm)
}

func (ctrl *RobotsController) Store(c *gin.Context) {
	request := requests.RobotRequest{}
	if ok := requests.Validate(c, &request, requests.RobotSave); !ok {
		return
	}

	rm := robot.Robot{
		Name:        request.Name,
		AccessToken: request.AccessToken,
		ProjectID:   request.ProjectID,
	}
	rm.Create()
	if rm.ID > 0 {
		response.Created(c, rm)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *RobotsController) Update(c *gin.Context) {

	robotModel := robot.Get(c.Param("id"))
	oldProjectID := robotModel.ProjectID
	if robotModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.RobotRequest{}
	if ok := requests.Validate(c, &request, requests.RobotSave); !ok {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	robotModel.Name = request.Name
	robotModel.AccessToken = request.AccessToken
	robotModel.ProjectID = request.ProjectID
	if ok := robotModel.Save(oldProjectID); !ok {
		response.Abort500(c, "更新失败，请稍后尝试~")
		return
	}

	response.Data(c, robotModel)
}

func (ctrl *RobotsController) Delete(c *gin.Context) {

	robotModel := robot.Get(c.Param("id"))
	if robotModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if !robotModel.Delete() {
		response.Abort500(c, "删除失败，请稍后尝试~")
		return
	}

	response.Success(c)
}
