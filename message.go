package alipay

import (
	"encoding/json"
	"errors"
)

// AlipayOpenAppMiniTemplateMessageSend 小程序发送模板消息
// 文档地址: https://opendocs.alipay.com/mini/02cth2
func (c *Client) AlipayOpenAppMiniTemplateMessageSend(bm BodyMap) (aliRsp *AliPayOpenAppMiniTemplateMessageSendResponse, err error) {
	if err = bm.CheckEmptyError("to_user_id", "user_template_id", "page", "data"); err != nil {
		return nil, err
	}

	var bs []byte
	if bs, err = c.doAlipayRequest(HttpPostMethod, bm, "alipay.open.app.mini.templatemessage.send", ""); err != nil {
		return nil, err
	}

	aliRsp = new(AliPayOpenAppMiniTemplateMessageSendResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}

	if err = c.AutoVerifySign(bs, aliRsp.AlipayCertSn, aliRsp.Sign); err != nil {
		return nil, err
	}

	if aliRsp.Response.ErrorResponse != nil && aliRsp.Response.ErrorResponse.Code.IsSuccess() == false {
		return aliRsp, errors.New(aliRsp.Response.ErrorResponse.Error())
	}

	return aliRsp, nil
}
