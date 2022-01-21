//Package original 模型
package original

import (
	"github.com/spf13/cast"
	"sentry-white-go/app/handlers/oss"
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

func (original *Original) Create() {
	originals := All()

	// 获取最大 ID
	var maxId uint64 = 0
	for _, o := range originals {
		if o.ID > maxId {
			maxId = o.ID
		}
	}

	// 给予当前模型 ID
	original.ID = maxId + 1
	originals = append(originals, *original)

	oss.Upload(ApiPath, originals)
	oss.Upload(ApiPath+"/"+cast.ToString(original.ID), original)
}

func (original *Original) Save() (rowsAffected int64) {
	originals := All()
	for i, o := range originals {
		if o.ID == original.ID {
			originals[i].Name = original.Name
			originals[i].AccountName = original.AccountName
			originals[i].Password = original.Password
		}
	}
	oss.Upload(ApiPath, originals)
	oss.Upload(ApiPath+"/"+cast.ToString(original.ID), original)
	return 1
}

func (original *Original) Delete() (rowsAffected int64) {
	originals := All()
	var index int
	for i, o := range originals {
		if o.ID == original.ID {
			index = i
		}
	}
	originals = append(originals[:index], originals[index+1:]...)
	oss.Upload(ApiPath, originals)
	oss.Delete(ApiPath + "/" + cast.ToString(original.ID))
	return 1
}
