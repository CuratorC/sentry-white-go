//Package responsible_person 模型
package responsible_person

import (
	"github.com/curatorc/cngf/app"
	"sentry-white-go/app/models"
)

const ApiPath = "api/v1/responsible_people"

type ResponsiblePerson struct {
	models.BaseModel

	// Put fields in here
	Name       string   `json:"name"`
	Phone      string   `json:"phone"`
	ProjectsID []uint64 `json:"projects_id"`

	models.CommonTimestampsField
}

type ResponsiblePeopleCollection struct {
	ResponsiblePeople []ResponsiblePerson `json:"responsible_people"`
	models.BaseCollection
}

func (responsiblePerson *ResponsiblePerson) Create() {
	// 全部负责人添加信息
	rpcl := All()
	// 给予当前模型 ID
	rpcl.MaxID += 1
	responsiblePerson.ID = rpcl.MaxID
	responsiblePerson.CreatedAt = app.Now()
	responsiblePerson.UpdatedAt = app.Now()
	rpcl.ResponsiblePeople = append(rpcl.ResponsiblePeople, *responsiblePerson)

	models.UploadToOss(ApiPath, rpcl, responsiblePerson)
}

func (responsiblePerson *ResponsiblePerson) Save() bool {
	responsiblePerson.UpdatedAt = app.Now()

	// 其次修改列表中的信息
	rpcl := All()
	for i, rp := range rpcl.ResponsiblePeople {
		if rp.ID == responsiblePerson.ID {
			rpcl.ResponsiblePeople[i].Name = responsiblePerson.Name
			rpcl.ResponsiblePeople[i].Phone = responsiblePerson.Phone

			rpcl.ResponsiblePeople[i].UpdatedAt = responsiblePerson.UpdatedAt
		}
	}
	models.UploadToOss(ApiPath, rpcl, responsiblePerson)
	return true
}

func (responsiblePerson *ResponsiblePerson) Delete() bool {

	rpcl := All()
	responsiblePerson.DeletedAt = app.Now()
	for i, rp := range rpcl.ResponsiblePeople {
		if rp.ID == responsiblePerson.ID {
			rpcl.ResponsiblePeople[i].DeletedAt = responsiblePerson.DeletedAt
		}
	}

	models.DeleteOnOss(ApiPath, rpcl, responsiblePerson)
	return true
}
