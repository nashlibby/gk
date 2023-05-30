package gk

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/http"
	"strings"
)

type AliyunOss struct {
	Config AliyunOssConfig
}

type AliyunOssConfig struct {
	Endpoint     string `json:"endpoint"`
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
	BucketName   string `json:"bucket_name"`
	IsCname      bool   `json:"is_cname"`
	Debug        bool   `json:"debug"`
	Prefix       string `json:"prefix"`
	CdnName      string `json:"cdn_name"`
}

func NewAliyunOss(config AliyunOssConfig) Oss {
	return &AliyunOss{
		Config: config,
	}
}

// 文件上传
func (a *AliyunOss) UploadFile(remotePath string, file []byte) (remoteUrl string, err error) {

	// 创建OSSClient实例。
	client, err := oss.New(a.Config.Endpoint, a.Config.AccessKey, a.Config.AccessSecret)
	if err != nil {
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(a.Config.BucketName)
	if err != nil {
		return
	}

	// 上传Byte数组
	err = bucket.PutObject(strings.TrimLeft(remotePath, "/"), bytes.NewReader(file))
	if err != nil {
		return remoteUrl, err
	}

	if a.Config.CdnName != "" {
		return "https://" + a.Config.CdnName + "/" + strings.TrimLeft(remotePath, "/"), nil
	} else {
		return "https://" + a.Config.BucketName + "." + a.Config.Endpoint + "/" + strings.TrimLeft(remotePath, "/"), nil
	}
}

// 绝对地址直接转换为阿里云地址
func (a *AliyunOss) TransformFile(remotePath, url string) (remoteUrl string, err error) {

	remoteFile, err := a.ReadRemoteFile(url)
	if err != nil {
		return
	}

	// 创建OSSClient实例。
	client, err := oss.New(a.Config.Endpoint, a.Config.AccessKey, a.Config.AccessSecret)
	if err != nil {
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(a.Config.BucketName)
	if err != nil {
		return
	}

	// 上传Byte数组
	err = bucket.PutObject(strings.TrimLeft(remotePath, "/"), bytes.NewReader(remoteFile))
	if err != nil {
		return remoteUrl, err
	}

	if a.Config.CdnName != "" {
		return "https://" + a.Config.CdnName + "/" + strings.TrimLeft(remotePath, "/"), nil
	} else {
		return "https://" + a.Config.BucketName + "." + a.Config.Endpoint + "/" + strings.TrimLeft(remotePath, "/"), nil
	}
}

// 远程文件读
func (a *AliyunOss) ReadRemoteFile(url string) (file []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
