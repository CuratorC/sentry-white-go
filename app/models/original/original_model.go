//Package original 模型
package original

import (
	"github.com/curatorc/cngf/app"
	"sentry-white-go/app/models"
	"sentry-white-go/app/models/project"
)

const ApiPath = "api/v1/originals"

type Original struct {
	models.BaseModel

	// Put fields in here
	Name        string             `json:"name"`
	AccountName string             `json:"account_name"`
	Password    string             `json:"password"`
	Projects    []*project.Project `json:"projects"`

	models.CommonTimestampsField
}

type OriginalsCollection struct {
	Originals []Original `json:"originals"`
	models.BaseCollection
}

func (original *Original) Create() {
	ocl := All()

	// 给予当前模型 ID
	ocl.MaxID += 1
	original.ID = ocl.MaxID
	original.CreatedAt = app.TimenowInTimezone()
	original.UpdatedAt = app.TimenowInTimezone()
	ocl.Originals = append(ocl.Originals, *original)

	models.UploadToOss(ApiPath, ocl, original)
}

func (original *Original) Save() (rowsAffected int64) {

	original.UpdatedAt = app.TimenowInTimezone()

	ocl := All()
	for i, o := range ocl.Originals {
		if o.ID == original.ID {
			ocl.Originals[i].Name = original.Name
			ocl.Originals[i].AccountName = original.AccountName
			ocl.Originals[i].Password = original.Password

			ocl.Originals[i].UpdatedAt = original.UpdatedAt
		}
	}
	models.UploadToOss(ApiPath, ocl, original)
	return 1
}

func (original *Original) Delete() (rowsAffected int64) {
	ocl := All()
	original.DeletedAt = app.TimenowInTimezone()
	for i, o := range ocl.Originals {
		if o.ID == original.ID {
			ocl.Originals[i].DeletedAt = original.DeletedAt
		}
	}

	models.DeleteOnOss(ApiPath, ocl, original)
	return 1
}
