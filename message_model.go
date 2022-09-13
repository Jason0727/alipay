package alipay

type AliPayOpenAppMiniTemplateMessageSendResponse struct {
	Response     *AliPayOpenAppMiniTemplateMessageSendData `json:"alipay_open_app_mini_templatemessage_send_response"`
	AlipayCertSn string                                    `json:"alipay_cert_sn,omitempty"`
	Sign         string                                    `json:"sign"`
}
type AliPayOpenAppMiniTemplateMessageSendData struct {
	*ErrorResponse
}
