package alipay

import (
	"encoding/json"
	"fmt"
	"strings"
)

type BodyMap map[string]interface{}

// Set 设置参数
func (bm BodyMap) Set(key string, value interface{}) BodyMap {
	bm[key] = value
	return bm
}

// SetBodyMap 回调函数设置 bm
func (bm BodyMap) SetBodyMap(key string, value func(b BodyMap)) BodyMap {
	_bm := make(BodyMap)
	value(_bm)
	bm[key] = _bm
	return bm
}

// 获取参数，同 GetString()
func (bm BodyMap) Get(key string) string {
	return bm.GetString(key)
}

// GetString 获取参数并转换成字符串
func (bm BodyMap) GetString(key string) string {
	if bm == nil {
		return EMPTY
	}
	value, ok := bm[key]
	if !ok {
		return EMPTY
	}
	v, ok := value.(string)
	if !ok {
		return bm.convertToString(value)
	}
	return v
}

// GetInterface 获取原始参数
func (bm BodyMap) GetInterface(key string) interface{} {
	if bm == nil {
		return nil
	}
	return bm[key]
}

// Remove 删除参数
func (bm BodyMap) Remove(key string) {
	delete(bm, key)
}

// Reset 置空 bm
func (bm BodyMap) Reset() {
	for k := range bm {
		delete(bm, k)
	}
}

// Marshal 将 bm 转换成 json
func (bm BodyMap) Marshal() (jb string) {
	bs, err := json.Marshal(bm)
	if err != nil {
		return ""
	}

	jb = string(bs)

	return jb
}

// Unmarshal 解析 bm 到 结构体或 map 中
func (bm BodyMap) Unmarshal(ptr interface{}) (err error) {
	bs, err := json.Marshal(bm)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}

// CheckEmptyError 校验指定key是否为空
func (bm BodyMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.GetString(k); v == EMPTY {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return fmt.Errorf("[%w], %v", MissParamErr, strings.Join(emptyKeys, ", "))
	}
	return nil
}

// convertToString 将任意类型转换成字符串
func (bm BodyMap) convertToString(v interface{}) (str string) {
	if v == nil {
		return EMPTY
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return EMPTY
	}
	str = string(bs)
	return
}
