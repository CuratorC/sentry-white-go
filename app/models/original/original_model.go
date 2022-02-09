//Package original 模型
package original

import (
	"github.com/curatorc/cngf/app"
	"sentry-white-go/app/models"
)

const ApiPath = "api/v1/originals"

type Original struct {
	models.BaseModel

	// Put fields in here
	Name        string     `json:"name"`
	AccountName string     `json:"account_name"`
	Password    string     `json:"password"`
	Projects    []*Project `json:"projects"`

	models.CommonTimestampsField
}

type Project struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	SubstituteName string `json:"substitute_name"`
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
	original.CreatedAt = app.Now()
	original.UpdatedAt = app.Now()
	ocl.Originals = append(ocl.Originals, *original)

	models.UploadToOss(ApiPath, ocl, original)
}

func (original *Original) Save() bool {

	original.UpdatedAt = app.Now()

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
	return true
}

func (original *Original) Delete() bool {
	ocl := All()
	original.DeletedAt = app.Now()
	for i, o := range ocl.Originals {
		if o.ID == original.ID {
			ocl.Originals[i].DeletedAt = original.DeletedAt
		}
	}

	models.DeleteOnOss(ApiPath, ocl, original)
	return true
}
