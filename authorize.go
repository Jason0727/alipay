package alipay

import (
	"encoding/json"
	"errors"
)

// AlipaySystemOauthToken 换取授权访问令牌
// 文档地址: https://opendocs.alipay.com/open/02xtla
func (c *Client) AlipaySystemOauthToken(bm BodyMap) (aliRsp *AlipaySystemOauthTokenResponse, err error) {
	if err = bm.CheckEmptyError("grant_type"); err != nil {
		return nil, err
	}
	if bm.Get("grant_type") == "authorization_code" {
		if err = bm.CheckEmptyError("code"); err != nil {
			return nil, err
		}
	} else if bm.Get("grant_type") == "refresh_token" {
		if err = bm.CheckEmptyError("refresh_token"); err != nil {
			return nil, err
		}
	}

	var bs []byte
	if bs, err = c.doAlipayRequest(HttpPostMethod, bm, "alipay.system.oauth.token", ""); err != nil {
		return nil, err
	}

	aliRsp = new(AlipaySystemOauthTokenResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}

	if err = c.AutoVerifySign(bs, aliRsp.AlipayCertSn, aliRsp.Sign); err != nil {
		return nil, err
	}

	if aliRsp.ErrorResponse != nil {
		return aliRsp, errors.New(aliRsp.ErrorResponse.Error())
	}

	return aliRsp, nil
}
