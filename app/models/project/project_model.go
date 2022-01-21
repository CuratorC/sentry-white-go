//Package project 模型
package project

import (
	"github.com/spf13/cast"
	"sentry-white-go/app/handlers/oss"
	"sentry-white-go/app/models"
	"sentry-white-go/app/models/responsible_person"
	"sentry-white-go/app/models/robot"
)

const ApiPath = "api/v1/projects"

type Project struct {
	models.BaseModel

	// Put fields in here
	Name              string                                  `json:"name"`
	SubstituteName    string                                  `json:"substitute_name"`
	Robot             *robot.Robot                            `json:"robot"`
	ResponsiblePeople []*responsible_person.ResponsiblePerson `json:"responsible_people"`

	models.CommonTimestampsField
}

func (project *Project) Create() {
	projects := All()

	// 获取最大 ID
	var maxId uint64 = 0
	for _, o := range projects {
		if o.ID > maxId {
			maxId = o.ID
		}
	}

	// 给予当前模型 ID
	project.ID = maxId + 1
	projects = append(projects, *project)

	oss.Upload(ApiPath, projects)
}

func (project *Project) Save() (rowsAffected int64) {
	projects := All()
	for i, o := range projects {
		if o.ID == project.ID {
			projects[i].Name = project.Name
			projects[i].SubstituteName = project.SubstituteName
		}
	}
	oss.Upload(ApiPath, projects)
	oss.Upload(ApiPath+"/"+cast.ToString(project.ID), project)
	return 1
}

func (project *Project) Delete() (rowsAffected int64) {
	projects := All()
	var index int
	for i, o := range projects {
		if o.ID == project.ID {
			index = i
		}
	}
	projects = append(projects[:index], projects[index+1:]...)
	oss.Upload(ApiPath, projects)
	oss.Delete(ApiPath + "/" + cast.ToString(project.ID))
	return 1
}
