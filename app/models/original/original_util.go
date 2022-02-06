package original

import (
	"encoding/json"
	"github.com/curatorc/cngf/cache"
	"github.com/curatorc/cngf/logger"
	"sentry-white-go/app/handlers/oss"
	"sentry-white-go/app/models"
)

func Get(idstr string) (original Original) {
	models.GetModelFromOSS(ApiPath+"/"+idstr, &original)
	return
}

func All() OriginalsCollection {
	wanted := OriginalsCollection{}
	cache.RememberObject("cache-key-"+ApiPath, 200, &wanted, func() {
		if oss.IsExist(ApiPath) {
			response := oss.Get(oss.SignURL(ApiPath))
			err := json.Unmarshal([]byte(response), &wanted)
			logger.LogIf(err)
		}
	})

	return wanted
}
