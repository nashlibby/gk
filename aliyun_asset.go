package gk

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"net/http"
	"strings"
)

type AliyunAsset struct {
	Endpoint     string `json:"endpoint"`
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
	BucketName   string `json:"bucket_name"`
}

func NewAliyunAsset(endpoint, accessKey, accessSecret, bucketName string) *AliyunAsset {
	return &AliyunAsset{
		Endpoint:     endpoint,
		AccessKey:    accessKey,
		AccessSecret: accessSecret,
		BucketName:   bucketName,
	}
}

// 本地上传
func (a *AliyunAsset) Upload(path string, file []byte) (remoteUrl string, err error) {

	// 创建OSSClient实例。
	client, err := oss.New(a.Endpoint, a.AccessKey, a.AccessSecret)
	if err != nil {
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(a.BucketName)
	if err != nil {
		return
	}

	// 上传Byte数组
	err = bucket.PutObject(strings.TrimLeft(path, "/"), bytes.NewReader(file))
	if err != nil {
		return remoteUrl, err
	}

	return "https://" + a.BucketName + "." + a.Endpoint + "/" + strings.TrimLeft(path, "/"), nil
}

// 绝对地址直接转换为阿里云地址
func (a *AliyunAsset) Transform(path, url string) (remoteUrl string, err error) {

	remoteFile, err := a.ReadRemoteFile(url)
	if err != nil {
		return
	}

	// 创建OSSClient实例。
	client, err := oss.New(a.Endpoint, a.AccessKey, a.AccessSecret)
	if err != nil {
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(a.BucketName)
	if err != nil {
		return
	}

	// 上传Byte数组
	err = bucket.PutObject(strings.TrimLeft(path, "/"), bytes.NewReader(remoteFile))
	if err != nil {
		return remoteUrl, err
	}

	return "https://" + a.BucketName + "." + a.Endpoint + "/" + strings.TrimLeft(path, "/"), nil
}

// 远程文件读
func (a *AliyunAsset) ReadRemoteFile(url string) (file []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
