package alipay

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-pay/gopay"
)

// 网站支付
func WebPageAlipay() (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", strconv.FormatInt(time.Now().Unix(), 10)).
		Set("total_amount", "0.01").
		Set("subject", "测试电脑网站支付").
		Set("product_code", "FAST_INSTANT_TRADE_PAY").
		Set("appid", appid)

	url, err := client.TradePagePay(context.TODO(), bm)
	if err != nil {
		log.Println("WebPageAlipay 网站支付下单出错", err)
		return "", err
	}

	log.Println("网站支付、客户端直接跳转即可：", url)
	return url, nil
}

// 移动端支付
func MobileAlipay() (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", strconv.FormatInt(time.Now().Unix(), 10)).
		Set("total_amount", "0.01").
		Set("subject", "测试手机网站支付").
		Set("product_code", "FAST_INSTANT_TRADE_PAY").
		Set("appid", appid)

	url, err := client.TradeWapPay(context.TODO(), bm)
	if err != nil {
		log.Println("MobileAlipay 手机网站支付下单出错", err)
		return "", err
	}

	log.Println("手机网站支付、客户端直接跳转即可：", url)
	return url, nil
}

// 商家码生成二维码
func ErweimeAlipay() (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", strconv.FormatInt(time.Now().Unix(), 10)).
		Set("total_amount", "0.01").
		Set("subject", "当面付扫商家二维码付款").
		Set("product_code", "FACE_TO_FACE_PAYMENT")

	aliRsp, err := client.TradePrecreate(context.TODO(), bm)
	if err != nil {
		log.Println("ErweimeAlipay 当面付支付下单出错", err)
		return "", err
	}

	log.Println("当面付、客户端需生成临时二维码、用户通过支付宝扫码支付：", aliRsp.Response.QrCode)
	return aliRsp.Response.QrCode, nil
}
