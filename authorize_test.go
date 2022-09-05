package alipay_test

import (
	"fmt"
	"github.com/Jason0727/alipay"
	"testing"
)

func TestClient_AlipaySystemOauthToken(t *testing.T) {
	t.Log("========== SystemOauthToken ==========")

	bm := make(alipay.BodyMap, 0)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", "72837c48a3024324b01ae9170dcaFX19")

	fmt.Println(client.AlipaySystemOauthToken(bm))
}
