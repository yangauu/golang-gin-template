package alipay

import (
	"log"
	"os"

	"github.com/go-pay/gopay/alipay"
)

var appid = "2021003126626960" // 应用ID
var PrivateKey = ""            // 应用私钥，支持PKCS1和PKCS8
var client *alipay.Client      // 支付客户端

func init() {
	exPath, _ := os.Getwd()
	pem, err := os.ReadFile(exPath + "/app/pay/alipay/cert/alipayRootCertKey.pem")
	if err != nil {
		log.Println("读取出错")
	}
	PrivateKey = string(pem)

	client, err = alipay.NewClient(appid, PrivateKey, true) // true/false 是否是正式环境
	if err != nil {
		log.Println("支付宝初始化出错")
		return
	}

	err = client.SetCertSnByPath(exPath+"/app/pay/alipay/cert/appCertPublicKey_2021003126626960.crt", exPath+"/app/pay/alipay/cert/alipayRootCert.crt", exPath+"/app/pay/alipay/cert/appCertPublicKey_2021003126626960.crt")
	if err != nil {
		log.Println("支付宝初始化加载证书出错")
	}
}
