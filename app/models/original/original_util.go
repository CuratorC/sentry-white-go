package original

import (
	"encoding/json"
	"github.com/curatorc/cngf/logger"
	"github.com/spf13/cast"
	"sentry-white-go/app/handlers/oss"
)

func Get(idstr string) (original Original) {
	originals := All()
	for _, o := range originals {
		if o.ID == cast.ToUint64(idstr) {
			original = o
		}
	}
	return
}

func All() (originals []Original) {
	if oss.IsExist(ApiPath) {
		response := oss.Get(oss.SignURL(ApiPath))
		err := json.Unmarshal([]byte(response), &originals)
		logger.LogIf(err)
	}
	return originals
}
