package oss

import (
	"bytes"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/curatorc/cngf/helpers"
	"github.com/curatorc/cngf/logger"
	"net/http"
	"sentry-white-go/app/models/local_storage"
	"strings"
)

// 获取 bucket 对象
func getBucket() (b *oss.Bucket) {
	lsm := local_storage.Get()

	c, err := oss.New(
		lsm.AliyunAccessOSSEndpoint,
		lsm.AliyunAccessID,
		lsm.AliyunAccessSecret,
		oss.Timeout(10, 120),
	)
	logger.LogIf(err)
	c.GetBucketStat(lsm.AliyunAccessOSSBucket)
	b, err = c.Bucket(lsm.AliyunAccessOSSBucket)
	logger.LogIf(err)
	return
}

// Upload 从字符串中上传文件
func Upload(fileName string, fileContent interface{}) {
	s, err := json.Marshal(&fileContent)
	logger.LogIf(err)

	bucket := getBucket()
	err = bucket.PutObject(fileName, strings.NewReader(string(s)))
	logger.LogIf(err)
}

// Delete 删除文件
func Delete(fileName string) {
	bucket := getBucket()
	err := bucket.DeleteObject(fileName)
	logger.LogIf(err)
}

// SignURL 获取签名地址
func SignURL(fileName string) (signedURL string) {
	bucket := getBucket()

	// 带可选参数的签名直传。请确保设置的ContentType值与在前端使用时设置的ContentType值一致。
	/*options := []oss.Option{
		oss.ContentType("application/json"),
	}*/

	signedURL, err := bucket.SignURL(fileName, oss.HTTPGet, 600)
	// signedURL, err := bucket.SignURL(fileName, oss.HTTPPut, 600, options...)
	logger.LogIf(err)
	return
}

// IsExist 判断文件是否存在
func IsExist(fileName string) (isExist bool) {
	bucket := getBucket()
	isExist, err := bucket.IsObjectExist(fileName)
	logger.LogIf(err)
	return isExist
}

// BucketConnectionSuccess 判断OSS权限是否通过
func BucketConnectionSuccess() (success bool) {
	lsm := local_storage.Get()
	c, err := oss.New(
		lsm.AliyunAccessOSSEndpoint,
		lsm.AliyunAccessID,
		lsm.AliyunAccessSecret,
		oss.Timeout(10, 120),
	)
	logger.LogIf(err)
	stat, err := c.GetBucketInfo(lsm.AliyunAccessOSSBucket)
	if err != nil || helpers.Empty(stat.BucketInfo.Name) {
		return false
	}
	return true
}

// Get 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	response, _ := http.Get(url)
	// response.Body类型为io.ReadCloser
	//fmt.Printf(response.Body)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(response.Body)
	if err != nil {
		return ""
	}

	return buf.String()
}
