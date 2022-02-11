//Package project 模型
package project

import (
	"github.com/curatorc/cngf/app"
	"github.com/curatorc/cngf/helpers"
	"github.com/curatorc/cngf/logger"
	"github.com/spf13/cast"
	"sentry-white-go/app/models"
	"sentry-white-go/app/models/original"
	"sentry-white-go/app/models/pivots"
)

const ApiPath = "api/v1/projects"

type Project struct {
	models.BaseModel

	// Put fields in here
	Name                string   `json:"name"`
	SubstituteName      string   `json:"substitute_name"`
	Robot               *Robot   `json:"robot"`
	ResponsiblePeopleID []uint64 `json:"responsible_people_id"`
	OriginalID          uint64   `json:"original_id"`

	models.CommonTimestampsField
}

type Robot struct {
	ID uint64 `json:"id"`
}

type ProjectsCollection struct {
	Projects []Project `json:"projects"`
	models.BaseCollection
}

func (project *Project) Create() {

	// 全部项目添加信息
	pcl := All()
	// 给予当前模型 ID
	pcl.MaxID += 1
	project.ID = pcl.MaxID
	project.CreatedAt = app.Now()
	project.UpdatedAt = app.Now()
	pcl.Projects = append(pcl.Projects, *project)

	models.UploadToOss(ApiPath, pcl, project)

	// 指定 Original 分配信息
	addToOriginal(project)
}

func (project *Project) Save(oldOriginalID uint64) bool {

	project.UpdatedAt = app.Now()

	// 首先修改父级中的信息
	if oldOriginalID == project.OriginalID { // 当组织ID没有改变时
		updateOnOriginal(project)
	} else { // 若组织 ID 改变了
		// 先从旧组织中删除
		deleteOnOriginal(project, oldOriginalID)

		// 然后添加到新组织内
		addToOriginal(project)

	}

	// 其次修改列表中的信息
	pcl := All()
	for i, p := range pcl.Projects {
		if p.ID == project.ID {
			pcl.Projects[i].Name = project.Name
			pcl.Projects[i].SubstituteName = project.SubstituteName
			pcl.Projects[i].OriginalID = project.OriginalID

			pcl.Projects[i].UpdatedAt = project.UpdatedAt
		}
	}
	models.UploadToOss(ApiPath, pcl, project)
	return true

}

func (project *Project) Delete() bool {

	// 获取所有负责人
	for _, rpID := range project.ResponsiblePeopleID {
		pivots.DeleteProjectFromResponsiblePerson(project.ID, rpID)
	}

	deleteOnOriginal(project, project.OriginalID)
	pcl := All()
	project.DeletedAt = app.Now()
	for i, o := range pcl.Projects {
		if o.ID == project.ID {
			pcl.Projects[i].DeletedAt = project.DeletedAt
		}
	}

	models.DeleteOnOss(ApiPath, pcl, project)
	return true
}

// deleteOnOriginal 从组织内删除项目
func deleteOnOriginal(project *Project, originalId uint64) bool {
	o := original.Get(cast.ToString(originalId))
	index := -1
	for i, p := range o.Projects {
		if p.ID == project.ID {
			index = i
		}
	}
	if index == -1 {
		return false
	}
	o.Projects = append(o.Projects[:index], o.Projects[index+1:]...)
	o.Save()
	return true
}

// addToOriginal 从组织内删除项目
func addToOriginal(project *Project) bool {
	o := original.Get(cast.ToString(project.OriginalID))
	oProject := original.Project{
		ID:             project.ID,
		Name:           project.Name,
		SubstituteName: project.SubstituteName,
	}
	o.Projects = append(o.Projects, &oProject)
	o.Save()
	return true
}

// updateOnOriginal 更新组织内的信息s
func updateOnOriginal(project *Project) bool {
	o := original.Get(cast.ToString(project.OriginalID))
	for i, p := range o.Projects {
		if p.ID == project.ID {
			o.Projects[i].Name = project.Name
		}
	}
	o.Save()
	return true
}

func DeleteResponsiblePersonFromProject(responsiblePersonID uint64, rpProjectsID []uint64) {

	for _, pID := range rpProjectsID {
		p := Get(cast.ToString(pID))
		// 从项目中移除负责人
		var index int
		for i, rpID := range p.ResponsiblePeopleID {
			if rpID == responsiblePersonID {
				index = i
			}
		}

		if helpers.Empty(index) {
			logger.ErrorString("responsible_person delete error", "responsible_person no in project", cast.ToString(p.ID))
		} else {
			p.ResponsiblePeopleID = append(p.ResponsiblePeopleID[:index], p.ResponsiblePeopleID[index+1:]...)
			p.Save(p.OriginalID)
		}
	}

}
