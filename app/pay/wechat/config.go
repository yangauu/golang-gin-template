package wechat

import (
	"log"
	"os"

	"github.com/go-pay/gopay/wechat/v3"
)

type JspOrders struct {
	SpMchid     string `json:"sp_mchid"`     // 服务商户号，由微信支付生成并下发
	SubMchid    string `json:"sub_mchid"`    // 子商户的商户号，由微信支付生成并下发
	OutTradeno  string `json:"out_trade_no"` // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	SpAppid     string `json:"sp_appid"`     // body 由微信生成的应用ID，全局唯一。请求基础下单接口时请注意APPID的应用属性，例如公众号场景下，需使用应用属性为公众号的服务号APPID
	SubAppid    string `json:"sub_appid"`    // 由微信生成的应用ID，全局唯一。请求基础下单接口时请注意APPID的应用属性，例如公众号场景下，需使用应用属性为公众号的服务号APPID
	Description string `json:"description"`  // 商品描述
	NotifyUrl   string `json:"notify_url"`   // 通知URL必须为直接可访问的URL，不允许携带查询串，要求必须为https地址
	Amount      Amount `json:"amount"`       // 订单金额信息
	Payer       Payer  `json:"payer"`        // 支付者信息
}

type Amount struct {
	Total    int    `json:"total"`    // 订单总金额，单位为分
	Currency string `json:"currency"` // 人民币，境内商户号仅支持人民币
}

type Payer struct {
	SpOpenid string `json:"sp_openid"` // 用户在服务商appid下的唯一标识。 下单前需获取到用户的Openid
}

type PrepayId struct {
	PrepayId string `json:"prepay_id"` // 发起请求的商户（包括直连商户、服务商或渠道商）的商户号
}

var MchId = "1611126463"                                  // 商户ID 或者服务商模式的 sp_mchid
var SerialNo = "587A0BDAC92E1EFD7FA47C298A387028D725F896" // 商户证书的证书序列号
var APIv3Key = "fhsr4g9GHDg8Fh6fYGHhjrg9G7GjLtX8"         // 商户平台获取
var PrivateKey = ""                                       // 私钥 apiclient_key.pem 读取后的内容
var client *wechat.ClientV3                               // 支付客户端

func init() {
	exPath, _ := os.Getwd()
	pem, err := os.ReadFile(exPath + "/app/pay/wechat/cert/apiclient_key.pem")
	if err != nil {
		log.Println("读取出错")
	}
	PrivateKey = string(pem)

	client, err = wechat.NewClientV3(MchId, SerialNo, APIv3Key, string(PrivateKey))
	if err != nil {
		log.Println("微信支付初始化出错")
	}
}
