package gk

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliyunSms struct {
	Config AliyunSmsConfig
}

type AliyunSmsConfig struct {
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
}

func NewAliyunSms(config AliyunSmsConfig) Sms {
	return &AliyunSms{
		Config: config,
	}
}

// 发送短信
func (a *AliyunSms) Send(phone, sign, template, data string) (SmsResponse, error) {
	client, err := dysmsapi.NewClientWithAccessKey("ap-northeast-1", a.Config.AccessKey, a.Config.AccessSecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = sign
	request.TemplateCode = template
	request.TemplateParam = data

	res, err := client.SendSms(request)
	return SmsResponse{
		RequestId:  res.RequestId,
		ResponseId: res.BizId,
		Message:    res.Message,
		Code:       res.Code,
	}, err
}
