package alipay_test

import (
	"fmt"
	"github.com/Jason0727/alipay"
	"os"
)

var client *alipay.Client

func init() {
	var err error
	client, err = alipay.NewClient(alipay.AppID, alipay.PrivateKey, true)

	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		os.Exit(-1)
	}

	if err = client.SetCertSn(alipay.AppCertPublicKey, alipay.AlipayRootCert, alipay.AlipayCertPublicKeyRsa2); err != nil {
		fmt.Println("设置证书失败, 错误信息为", err)
		os.Exit(-1)
	}
}
