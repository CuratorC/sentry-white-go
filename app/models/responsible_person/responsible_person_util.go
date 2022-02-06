package responsible_person

import (
	"encoding/json"
	"github.com/curatorc/cngf/cache"
	"github.com/curatorc/cngf/logger"
	"sentry-white-go/app/handlers/oss"
	"sentry-white-go/app/models"
)

func Get(idstr string) (responsiblePerson ResponsiblePerson) {
	models.GetModelFromOSS(ApiPath+"/"+idstr, &responsiblePerson)
	return
}

func All() ResponsiblePeopleCollection {
	wanted := ResponsiblePeopleCollection{}
	cache.RememberObject("cache-key-"+ApiPath, 200, &wanted, func() {
		logger.DebugString("responsible_people", "all", "remote")
		if oss.IsExist(ApiPath) {
			response := oss.Get(oss.SignURL(ApiPath))
			err := json.Unmarshal([]byte(response), &wanted)
			logger.LogIf(err)
		}
	})

	return wanted
}
