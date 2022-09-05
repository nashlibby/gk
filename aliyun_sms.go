package gk

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliyunSms struct {
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
}

type Response struct {
	RequestId  string `json:"request_id"`
	ResponseId string `json:"response_id"`
	Message    string `json:"message"`
	Code       string `json:"code"`
}

func NewAliyunSms(accessKey, accessSecret string) *AliyunSms {
	return &AliyunSms{
		AccessKey:    accessKey,
		AccessSecret: accessSecret,
	}
}

// 发送短信
func (a *AliyunSms) Send(phone, sign, template, data string) (Response, error) {
	client, err := dysmsapi.NewClientWithAccessKey("ap-northeast-1", a.AccessKey, a.AccessSecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = sign
	request.TemplateCode = template
	request.TemplateParam = data

	res, err := client.SendSms(request)
	return Response{
		RequestId:  res.RequestId,
		ResponseId: res.BizId,
		Message:    res.Message,
		Code:       res.Code,
	}, err
}
