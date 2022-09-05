package alipay_test

import (
	"fmt"
	"github.com/Jason0727/alipay"
	"testing"
)

func TestClient_AlipayOpenAppMiniTemplateMessageSend(t *testing.T) {
	t.Log("========== SystemOauthToken ==========")
	bm := make(alipay.BodyMap, 0)
	bm.Set("to_user_id", "2088802504231191")
	bm.Set("user_template_id", "777f979b2f234c908b2e32de14857bec")
	bm.Set("page", "/pages/index/index")
	bm.Set("data", `{"keyword1":{"value":"旧衣回收"},"keyword2":{"value":"预约成功"},"keyword3":{"value":"感谢您帮地球减负~"}}`)

	fmt.Println(client.AlipayOpenAppMiniTemplateMessageSend(bm))
}
