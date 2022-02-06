package pivots

import (
	"github.com/curatorc/cngf/helpers"
	"github.com/curatorc/cngf/logger"
	"github.com/spf13/cast"
	"sentry-white-go/app/models/responsible_person"
)

func BindProjectResponsiblePerson(responsiblePersonID uint64, projectID uint64, pResponsiblePeopleID []uint64) {
	// 1. 判断是否存在
	if !helpers.Uint64InSlice(responsiblePersonID, pResponsiblePeopleID) {
		pResponsiblePeopleID = append(pResponsiblePeopleID, responsiblePersonID)
	}
	rp := responsible_person.Get(cast.ToString(responsiblePersonID))
	// 1. 判断是否存在
	if !helpers.Uint64InSlice(projectID, rp.ProjectsID) {
		rp.ProjectsID = append(rp.ProjectsID, projectID)
	}
	rp.Save()
}

func UnBindProjectResponsiblePerson(responsiblePersonID uint64, projectID uint64, pResponsiblePeopleID []uint64) {
	rp := responsible_person.Get(cast.ToString(responsiblePersonID))
	// 1. 判断是否存在
	if helpers.Uint64InSlice(responsiblePersonID, pResponsiblePeopleID) {
		var indexRp int
		for i, rpID := range pResponsiblePeopleID {
			if rpID == responsiblePersonID {
				indexRp = i
			}
		}

		if !helpers.Empty(indexRp) {
			pResponsiblePeopleID = append(pResponsiblePeopleID[:indexRp], pResponsiblePeopleID[indexRp+1:]...)
		}
	}
	// 1. 判断是否存在
	if helpers.Uint64InSlice(projectID, rp.ProjectsID) {
		var indexP int
		for i, pID := range rp.ProjectsID {
			if pID == projectID {
				indexP = i
			}
		}

		if !helpers.Empty(indexP) {
			rp.ProjectsID = append(rp.ProjectsID[:indexP], rp.ProjectsID[indexP+1:]...)
		}
	}
	rp.Save()
}

func DeleteProjectFromResponsiblePerson(projectID uint64, responsiblePersonID uint64) {

	rp := responsible_person.Get(cast.ToString(responsiblePersonID))
	// 从负责人中移除项目
	var index int
	for i, pID := range rp.ProjectsID {
		if pID == projectID {
			index = i
		}
	}

	if helpers.Empty(index) {
		logger.ErrorString("project delete error", "project_id no in responsible_person", cast.ToString(responsiblePersonID))
	} else {
		rp.ProjectsID = append(rp.ProjectsID[:index], rp.ProjectsID[index+1:]...)
		rp.Save()
	}

}
