package gk

// 短信响应
type SmsResponse struct {
	RequestId  string `json:"request_id"`
	ResponseId string `json:"response_id"`
	Message    string `json:"message"`
	Code       string `json:"code"`
}

type Sms interface {
	// 发送短信
	Send(phone, sign, template, data string) (SmsResponse, error)
}
