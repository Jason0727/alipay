package alipay_test

import (
	"fmt"
	"github.com/Jason0727/alipay"
	"testing"
)

func TestClient_AlipaySystemOauthToken(t *testing.T) {
	t.Log("========== AlipaySystemOauthToken ==========")

	bm := make(alipay.BodyMap, 0)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", "1aacc8159cb94afebb4b9c4ef3f1FX19")

	fmt.Println(client.AlipaySystemOauthToken(bm))
}

func TestClient_AlipayUserInfoShare(t *testing.T) {
	t.Log("========== AlipayUserInfoShare ==========")

	authToken := "authusrB2911e3f598ba436abe95d30860198X19"
	fmt.Println(client.AlipayUserInfoShare(authToken))

}