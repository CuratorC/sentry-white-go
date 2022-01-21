//Package robot 模型
package robot

import (
	"github.com/curatorc/cngf/database"
	"sentry-white-go/app/models"
)

type Robot struct {
	models.BaseModel

	// Put fields in here
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`

	models.CommonTimestampsField
}

func (robot *Robot) Create() {
	database.DB.Create(&robot)
}

func (robot *Robot) Save() (rowsAffected int64) {
	result := database.DB.Save(&robot)
	return result.RowsAffected
}

func (robot *Robot) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&robot)
	return result.RowsAffected
}
