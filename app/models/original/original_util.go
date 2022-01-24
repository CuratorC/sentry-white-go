package original

import (
	"encoding/json"
	"github.com/curatorc/cngf/cache"
	"github.com/curatorc/cngf/logger"
	"sentry-white-go/app/handlers/oss"
)

func Get(idstr string) (original Original) {
	path := ApiPath + "/" + idstr
	if oss.IsExist(path) {
		response := oss.Get(oss.SignURL(path))
		err := json.Unmarshal([]byte(response), &original)
		logger.LogIf(err)
	}
	return
}

var cacheKey = "cache-key-" + ApiPath

func All() OriginalsCollection {
	return cache.Remember(cacheKey, 200, func() interface{} {
		var ocl OriginalsCollection
		if oss.IsExist(ApiPath) {
			response := oss.Get(oss.SignURL(ApiPath))
			err := json.Unmarshal([]byte(response), &ocl)
			logger.LogIf(err)
		}
		logger.WarnJSON("originals", "all", ocl)
		return ocl
	}).(OriginalsCollection)
}
