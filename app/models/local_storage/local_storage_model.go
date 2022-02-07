package local_storage

import (
	"encoding/json"
	"github.com/curatorc/cngf/logger"
	"io/ioutil"
	"os"
)

type LocalStorage struct {
	AliyunAccessID          string `json:"aliyun_access_id"`
	AliyunAccessSecret      string `json:"aliyun_access_secret"`
	AliyunAccessOSSEndpoint string `json:"aliyun_access_oss_endpoint"`
	AliyunAccessOSSBucket   string `json:"aliyun_access_oss_bucket"`
}

const FileName = `storage/models/local_storage.model`

func (storage *LocalStorage) Save() bool {
	err := os.MkdirAll(`storage/models/`, os.ModePerm)
	s, err := json.Marshal(&storage)
	logger.LogIf(err)
	// 保存到文件
	err = ioutil.WriteFile(FileName, s, 0666)
	logger.LogIf(err)
	return true
}

func Get() (storage LocalStorage) {
	_, err := os.Stat(FileName)
	if os.IsNotExist(err) {
		return
	}
	s, err := os.ReadFile(FileName)
	logger.LogIf(err)

	err = json.Unmarshal(s, &storage)
	logger.LogIf(err)
	return
}
