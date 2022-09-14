package alipay

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	xaes "github.com/Jason0727/alipay/pkg/aes"
	"reflect"
)

// DecryptOpenDataToStruct 解密支付宝开放数据到 结构体
//	encryptedData:包括敏感数据在内的完整用户信息的加密数据
//	secretKey:AES密钥，支付宝管理平台配置
//	beanPtr:需要解析到的结构体指针
//	文档：https://opendocs.alipay.com/mini/introduce/aes
//	文档：https://opendocs.alipay.com/open/common/104567
func DecryptOpenDataToStruct(encryptedData, secretKey string, beanPtr interface{}) (err error) {
	if encryptedData == EMPTY || secretKey == EMPTY {
		return errors.New("encryptedData or secretKey is null")
	}
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入参数类型必须是以指针形式")
	}
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	var (
		block      cipher.Block
		blockMode  cipher.BlockMode
		originData []byte
	)
	aesKey, _ := base64.StdEncoding.DecodeString(secretKey)
	ivKey := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)
	if block, err = aes.NewCipher(aesKey); err != nil {
		return fmt.Errorf("aes.NewCipher：%w", err)
	}
	if len(secretData)%len(aesKey) != 0 {
		return errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = xaes.PKCS5UnPadding(originData)
	}
	if err = json.Unmarshal(originData, beanPtr); err != nil {
		return fmt.Errorf("json.Unmarshal(%s)：%w", string(originData), err)
	}
	return nil
}

// DecryptOpenDataToBodyMap 解密支付宝开放数据到 BodyMap
//	encryptedData:包括敏感数据在内的完整用户信息的加密数据
//	secretKey:AES密钥，支付宝管理平台配置
//	文档：https://opendocs.alipay.com/mini/introduce/aes
//	文档：https://opendocs.alipay.com/open/common/104567
func DecryptOpenDataToBodyMap(encryptedData, secretKey string) (bm BodyMap, err error) {
	if encryptedData == EMPTY || secretKey == EMPTY {
		return nil, errors.New("encryptedData or secretKey is null")
	}
	var (
		aesKey, originData []byte
		ivKey              = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		block              cipher.Block
		blockMode          cipher.BlockMode
	)
	aesKey, _ = base64.StdEncoding.DecodeString(secretKey)
	secretData, _ := base64.StdEncoding.DecodeString(encryptedData)
	if block, err = aes.NewCipher(aesKey); err != nil {
		return nil, fmt.Errorf("aes.NewCipher：%w", err)
	}
	if len(secretData)%len(aesKey) != 0 {
		return nil, errors.New("encryptedData is error")
	}
	blockMode = cipher.NewCBCDecrypter(block, ivKey)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = xaes.PKCS5UnPadding(originData)
	}
	bm = make(BodyMap)
	if err = json.Unmarshal(originData, &bm); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(originData), err)
	}
	return
}
