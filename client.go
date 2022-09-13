package alipay

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/smartwalle/crypto4go"
)

type Client struct {
	appID            string          // 应用ID
	apiDomain        string          // 支付宝网关
	Client           *http.Client    // http 客户端
	location         *time.Location  // 时区
	privateKey       *rsa.PrivateKey // 应用私钥
	appCertSN        string          // 应用证书(根据内容生成SN)
	aliPayRootCertSN string          // 根证书(根据内容生成SN)
	aliPublicCertSN  string          // 支付宝公钥证书(根据内容生成SN)
	aliPayPublicKey  *rsa.PublicKey  // 支付宝证书公钥内容(alipayCertPublicKey_RSA2.crt)
}

// NewClient 实例化支付宝客户端
// appId: 应用ID privateKey: 应用私钥 isProd: 是否未生产环境
func NewClient(appId, privateKey string, isProd bool) (*Client, error) {
	if appId == EMPTY || privateKey == EMPTY {
		return nil, ErrInitParamsMissed
	}

	client := &Client{
		appID: appId,
	}

	if isProd {
		client.apiDomain = ProdURL
	} else {
		client.apiDomain = SandboxURL
	}

	client.Client = http.DefaultClient

	client.location = time.Local

	priKey, err := crypto4go.ParsePKCS1PrivateKey(crypto4go.FormatPKCS1PrivateKey(privateKey))
	if err != nil {
		priKey, err = crypto4go.ParsePKCS8PrivateKey(crypto4go.FormatPKCS8PrivateKey(privateKey))
		if err != nil {
			return nil, err
		}
	}
	client.privateKey = priKey

	return client, nil
}

// ============================================== 证书 ==================================
// SetCertSn 设置证书序列号
// 仅支持公钥证书模式
func (c *Client) SetCertSn(appCertPublicKeyContent, aliPayRootCertContent, aliPayPublicKeyRsaContent string) (err error) {
	// 加载应用公钥证书
	if err = c.loadAppCertPublicKey(appCertPublicKeyContent); err != nil {
		return err
	}

	// 加载支付宝根证书
	if err = c.loadAliPayRootCert(aliPayRootCertContent); err != nil {
		return err
	}

	// 加载支付宝公钥证书
	if err = c.loadAliPayPublicKeyRsaCert(aliPayPublicKeyRsaContent); err != nil {
		return err
	}

	return nil
}

// loadAppCertPublicKey 加载应用证书（将内容转换成SN）
// appCertPublicKey_2016073100129537.crt 文件内容
func (c *Client) loadAppCertPublicKey(s string) error {
	cert, err := crypto4go.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}

	c.appCertSN = c.getCertSN(cert)

	return nil
}

// loadAliPayRootCert 加载支付宝根证书（将内容转换成SN）
// alipayRootCert.crt 文件内容
func (c *Client) loadAliPayRootCert(s string) error {
	var certStrList = strings.Split(s, CertificateEnd)

	var certSNList = make([]string, 0, len(certStrList))

	for _, certStr := range certStrList {
		certStr = certStr + CertificateEnd

		var cert, _ = crypto4go.ParseCertificate([]byte(certStr))
		if cert != nil && (cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA1WithRSA) {
			certSNList = append(certSNList, c.getCertSN(cert))
		}
	}

	c.aliPayRootCertSN = strings.Join(certSNList, "_")

	return nil
}

// loadAliPayPublicKeyRsaCert 加载支付宝公钥证书（将内容转换成SN）
// alipayCertPublicKey_RSA2.crt 文件内容
func (c *Client) loadAliPayPublicKeyRsaCert(s string) error {
	cert, err := crypto4go.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return nil
	}

	c.aliPublicCertSN = c.getCertSN(cert)

	c.aliPayPublicKey = key

	return nil
}

// getCertSN 获取证书的十六进制编码
func (c *Client) getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))

	return hex.EncodeToString(value[:])
}

// doAlipayRequest 发送请求
func (c *Client) doAlipayRequest(httpMethod string, bm BodyMap, method string, authToken string) (bs []byte, err error) {
	var (
		bizContent string
		bodyBytes  []byte
	)

	// 设置请求参数集合
	if bm != nil {
		if bodyBytes, err = json.Marshal(bm); err != nil {
			return nil, err
		}

		bizContent = string(bodyBytes)
	}

	// 设置公共请求参数
	var p = url.Values{}
	p.Add("app_id", c.appID)
	p.Add("method", method)
	p.Add("format", Format)
	p.Add("charset", Charset)
	p.Add("sign_type", SignType)
	p.Add("timestamp", time.Now().In(c.location).Format(TimeFormat))
	p.Add("version", Version)
	p.Add("app_cert_sn", c.appCertSN)
	p.Add("alipay_root_cert_sn", c.aliPayRootCertSN)
	if authToken != "" {
		p.Add("auth_token", authToken)
	}
	p.Add("biz_content", bizContent)

	// 特殊情况处理
	if method == "alipay.system.oauth.token" { // 没有biz_content
		p.Add("grant_type", bm.Get("grant_type"))
		p.Add("code", bm.Get("code"))
		p.Add("refresh_token", bm.Get("refresh_token"))
	}

	// 签名
	var sign string
	if sign, err = c.signWithPKCS1v15(p, c.privateKey, crypto.SHA256); err != nil {
		return nil, err
	}
	p.Add("sign", sign)

	// 发起请求
	buf := strings.NewReader(p.Encode())
	var request *http.Request
	if request, err = http.NewRequest(httpMethod, c.apiDomain, buf); err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", ContentType)
	var response *http.Response
	if response, err = c.Client.Do(request); err != nil {
		return nil, err
	}
	if response != nil {
		defer func() {
			_ = response.Body.Close()
		}()
	}
	var data []byte
	if data, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	return data, nil
}
