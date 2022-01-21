//Package responsible_person 模型
package responsible_person

import (
	"github.com/curatorc/cngf/database"
	"sentry-white-go/app/models"
)

type ResponsiblePerson struct {
	models.BaseModel

	// Put fields in here
	Name  string `json:"name"`
	Phone string `json:"phone"`

	models.CommonTimestampsField
}

func (responsiblePerson *ResponsiblePerson) Create() {
	database.DB.Create(&responsiblePerson)
}

func (responsiblePerson *ResponsiblePerson) Save() (rowsAffected int64) {
	result := database.DB.Save(&responsiblePerson)
	return result.RowsAffected
}

func (responsiblePerson *ResponsiblePerson) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&responsiblePerson)
	return result.RowsAffected
}
