package alipay

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"github.com/smartwalle/crypto4go"
	"net/url"
	"sort"
	"strings"
)

// AutoVerifySign 自动验签
// bs: 响应完整字节数据 alipayCertSn:支付宝返回的证书序号 sign: 支付宝响应的签名
func (this *Client) AutoVerifySign(bs []byte, alipayCertSn, sign string) (err error) {
	var signData string
	if signData, err = this.getSignData(bs, alipayCertSn); err != nil {
		return err
	}

	return this.autoVerifySignByCert(sign, signData)
}

// getSignData 获取签名数据
// bs 响应完整字节数据 alipayCertSn 支付宝返回的证书序号
func (this *Client) getSignData(bs []byte, alipayCertSn string) (signData string, err error) {
	if alipayCertSn != this.aliPublicCertSN {
		return EMPTY, ErrAliCertSnNotMatch
	}

	str := string(bs)
	bsLen := len(str)

	indexStart := strings.Index(str, `_response":`)
	indexStart = indexStart + 11
	indexEnd := strings.Index(str, `,"alipay_cert_sn":`)
	if indexEnd > indexStart && bsLen > indexStart {
		return str[indexStart:indexEnd], nil
	}

	return EMPTY, ErrGetSignData
}

// autoVerifySignByCert 同步验签
// sign: 支付宝响应的签名 signData: 支付宝响应待签名的字符串
func (this *Client) autoVerifySignByCert(sign string, signData string) (err error) {
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	hashs := crypto.SHA256
	h := hashs.New()
	h.Write([]byte(signData))
	if err = rsa.VerifyPKCS1v15(this.aliPayPublicKey, hashs, h.Sum(nil), signBytes); err != nil {
		return ErrVerifySignature
	}

	return nil
}

// signWithPKCS1v15 签名
func (this *Client) signWithPKCS1v15(param url.Values, privateKey *rsa.PrivateKey, hash crypto.Hash) (s string, err error) {
	if param == nil {
		param = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var src = strings.Join(pList, "&")
	sig, err := crypto4go.RSASignWithKey([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}
