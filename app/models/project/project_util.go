package project

import (
    "encoding/json"
    "github.com/curatorc/cngf/logger"
    "github.com/spf13/cast"
    "sentry-white-go/app/handlers/oss"
)

func Get(idstr string) (project Project) {
	projects := All()
	for _, p := range projects {
		if p.ID == cast.ToUint64(idstr) {
			project = p
		}
	}
	return
}

func All() (projects []Project) {
	if oss.IsExist(ApiPath) {
		response := oss.Get(oss.SignURL(ApiPath))
		err := json.Unmarshal([]byte(response), &projects)
		logger.LogIf(err)
	}
	return projects
}
