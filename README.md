AliPay SDK for Golang

## 安装

##### 启用 Go module

```
go get github.com/Jason0727/alipay
```

```
import github.com/Jason0727/alipay
```


## 帮助

在集成的过程中有遇到问题，欢迎加 微信: 13585966246

## 版本如何初始化

**下面用到的 privateKey 需要特别注意一下，如果是通过“支付宝开发平台开发助手”创建的CSR文件，在 CSR 文件所在的目录下会生成相应的私钥文件，我们需要使用该私钥进行签名。**

```
var client, err = alipay.NewClient(appID, privateKey, true)
```

## 已实现接口

* **手机网站支付接口**

  alipay.trade.wap.pay - **TradeWapPay()**

## 集成流程

从[支付宝开放平台](https://open.alipay.com/)申请创建相关的应用，使用自己的支付宝账号登录即可。

#### 沙箱环境

支付宝开放平台为每一个应用提供了沙箱环境，供开发人员开发测试使用。

沙箱环境是独立的，每一个应用都会有一个商家账号和买家账号。

#### 应用信息配置

参考[官网文档](https://docs.open.alipay.com/200/105894) 进行应用的配置。

本 SDK 中的签名方法默认为 **RSA2**，采用支付宝提供的 [RSA签名&验签工具](https://docs.open.alipay.com/291/105971) 生成秘钥时，秘钥的格式必须为 **PKCS1**，秘钥长度推荐 **2048**。所以在支付宝管理后台请注意配置 **RSA2(SHA256)密钥**。

生成秘钥对之后，将公钥提供给支付宝（通过支付宝后台上传）对我们请求的数据进行签名验证，我们的代码中将使用私钥对请求数据签名。

请参考 [如何生成 RSA 密钥](https://docs.open.alipay.com/291/105971)。

#### 关于应用私钥 (privateKey)

应用私钥是我们通过工具生成的私钥，调用支付宝接口的时候，我们需要使用该私钥对参数进行签名。

#### 支持 RSA 签名及验证

默认采用的是 RSA2 签名，不再提供 RSA 的支持

## License

This project is licensed under the MIT License.
