package alipay

import (
	"fmt"
	"testing"
)

func TestClient_DecryptOpenDataToBodyMap(t *testing.T) {
	t.Log("========== DecryptOpenDataToBodyMap ==========")

	data := `SwUI+EFmW4gCuODaUzriF2lL+WKxKU3kYtwoAtsav6+BSaEsfHSZ7IREg+GpKbfeilvYA3hUQ1tfAW94iMi5s4r1DbUTasqCeU09ttRW8XcJy13tC+CpkzAAgWDodvSKAXm6iTuQDiFYHMGHq19RgswQqAlKUGrS2ivEYgsY75QKMU1JzZvdLF0LEXsGIbPIYeP+S7h22YCWQkDVSX5frxxRedxP64SoLuON5ZyP6SBErnu77zu7bXgszXT8RaQVxUW3ltXNqr3OqBmZwMLutrkR6//HzJKstCcwVvEGQFBT8c1cHNF/HCuvdiCo+7bAZ1Kot10zGQHtKFZtV7KlJT7BufFZ0tCTaYx4LgPG/Eg=`
	key := `8U1JP9eZcqUKurzaLnlSNg==`
	result, err := DecryptOpenDataToBodyMap(data, key)
	fmt.Println(err, result)
}

type DecryptOpenDataToStructRsp struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"subCode,omitempty"`
	SubMsg  string `json:"subMsg,omitempty"`
}

func TestClient_DecryptOpenDataToStruct(t *testing.T) {
	t.Log("========== DecryptOpenDataToStruct ==========")

	data := `SwUI+EFmW4gCuODaUzriF2lL+WKxKU3kYtwoAtsav6+BSaEsfHSZ7IREg+GpKbfeilvYA3hUQ1tfAW94iMi5s4r1DbUTasqCeU09ttRW8XcJy13tC+CpkzAAgWDodvSKAXm6iTuQDiFYHMGHq19RgswQqAlKUGrS2ivEYgsY75QKMU1JzZvdLF0LEXsGIbPIYeP+S7h22YCWQkDVSX5frxxRedxP64SoLuON5ZyP6SBErnu77zu7bXgszXT8RaQVxUW3ltXNqr3OqBmZwMLutrkR6//HzJKstCcwVvEGQFBT8c1cHNF/HCuvdiCo+7bAZ1Kot10zGQHtKFZtV7KlJT7BufFZ0tCTaYx4LgPG/Eg=`
	key := `8U1JP9eZcqUKurzaLnlSNg==`

	rsp := new(DecryptOpenDataToStructRsp)
	err := DecryptOpenDataToStruct(data, key, rsp)
	fmt.Println(err,rsp)
}
