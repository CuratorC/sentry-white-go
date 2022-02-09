//Package robot 模型
package robot

import (
	"github.com/curatorc/cngf/app"
	"github.com/curatorc/cngf/logger"
	"github.com/spf13/cast"
	"sentry-white-go/app/models"
	"sentry-white-go/app/models/project"
)

const ApiPath = "api/v1/robots"

type Robot struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
	ProjectID   uint64 `json:"project_id"`

	models.CommonTimestampsField
}

type RobotsCollection struct {
	Robots []Robot `json:"robots"`
	models.BaseCollection
}

func (robot *Robot) Create() {
	// 全部机器人添加信息
	rcl := All()
	// 给予当前模型 ID
	rcl.MaxID += 1
	robot.ID = rcl.MaxID
	robot.CreatedAt = app.Now()
	robot.UpdatedAt = app.Now()
	rcl.Robots = append(rcl.Robots, *robot)

	models.UploadToOss(ApiPath, rcl, robot)

	// 指定 Original 分配信息
	addToProject(robot)
}

func (robot *Robot) Save(oldProjectID uint64) bool {
	robot.UpdatedAt = app.Now()

	// 首先修改父级中的信息
	if oldProjectID == robot.ProjectID { // 当组织ID没有改变时
		updateOnProject(robot)
	} else { // 若组织 ID 改变了
		// 先从旧组织中删除
		deleteOnProject(oldProjectID)

		// 然后添加到新组织内
		addToProject(robot)

	}

	// 其次修改列表中的信息
	rcl := All()
	for i, r := range rcl.Robots {
		if r.ID == robot.ID {
			rcl.Robots[i].Name = robot.Name
			rcl.Robots[i].AccessToken = robot.AccessToken
			rcl.Robots[i].ProjectID = robot.ProjectID

			rcl.Robots[i].UpdatedAt = robot.UpdatedAt
		}
	}
	models.UploadToOss(ApiPath, rcl, robot)
	return true

}

func (robot *Robot) Delete() bool {
	deleteOnProject(robot.ProjectID)
	rcl := All()
	robot.DeletedAt = app.Now()
	for i, r := range rcl.Robots {
		if r.ID == robot.ID {
			rcl.Robots[i].DeletedAt = robot.DeletedAt
		}
	}

	models.DeleteOnOss(ApiPath, rcl, robot)
	return true
}

// deleteOnProject 从项目内删除机器人
func deleteOnProject(projectID uint64) bool {
	p := project.Get(cast.ToString(projectID))
	p.Robot = nil
	logger.DebugJSON("robot", "deleteOnProject", p)
	p.Save(p.OriginalID)
	return true
}

// addToProject 向项目内添加机器人
func addToProject(robot *Robot) bool {
	p := project.Get(cast.ToString(robot.ProjectID))
	p.Robot = &project.Robot{
		ID: robot.ID,
	}
	p.Save(p.OriginalID)
	return true
}

// updateOnProject 更新项目内的信息
func updateOnProject(robot *Robot) bool {
	p := project.Get(cast.ToString(robot.ProjectID))
	p.Robot = &project.Robot{
		ID: robot.ID,
	}
	p.Save(p.OriginalID)
	return true
}
