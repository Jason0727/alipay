package alipay

type AlipaySystemOauthTokenResponse struct {
	Response      *AlipaySystemOauthTokenData `json:"alipay_system_oauth_token_response,omitempty"`
	ErrorResponse *ErrorResponse              `json:"error_response,omitempty"`
	AlipayCertSn  string                      `json:"alipay_cert_sn,omitempty"`
	Sign          string                      `json:"sign,omitempty"`
}

type AlipaySystemOauthTokenData struct {
	AccessToken  string `json:"access_token,omitempty"`
	AlipayUserId string `json:"alipay_user_id,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	ReExpiresIn  int64  `json:"re_expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	UserId       string `json:"user_id,omitempty"`
}