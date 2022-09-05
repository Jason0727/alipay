package alipay

import (
	"errors"
	"fmt"
)

var (
	ErrInitParamsMissed  = errors.New("客户端初始化参数异常")
	MissParamErr         = errors.New("参数不能为空")
	ErrAliCertSnNotMatch = errors.New("当前使用的支付宝公钥证书SN与网关响应报文中的SN不匹配")
	ErrGetSignData       = errors.New("获取报文签名数据失败")
	ErrVerifySignature   = errors.New("验签签名失败")
)

const (
	CodeSuccess          Code = "10000" // 接口调用成功
	CodeUnknowError      Code = "20000" // 服务不可用
	CodeInvalidAuthToken Code = "20001" // 授权权限不足
	CodeMissingParam     Code = "40001" // 缺少必选参数
	CodeInvalidParam     Code = "40002" // 非法的参数
	CodeBusinessFailed   Code = "40004" // 业务处理失败
	CodePermissionDenied Code = "40006" // 权限不足
)

type Code string

func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}

type ErrorResponse struct {
	Code    Code   `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, e.Code, e.Msg, e.SubCode, e.SubMsg)
}
